// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"

	"storj.io/storj/internal/processgroup"
	"storj.io/storj/pkg/utils"
)

// Processes contains list of processes
type Processes struct {
	List []*Process
}

// Exec executes a command on all processes
func (processes *Processes) Exec(ctx context.Context, command string) error {
	var group errgroup.Group
	processes.Start(ctx, &group, command)
	return group.Wait()
}

// Start executes all processes using specified errgroup.Group
func (processes *Processes) Start(ctx context.Context, group *errgroup.Group, command string) {
	for _, p := range processes.List {
		process := p
		group.Go(func() error {
			return process.Exec(ctx, command)
		})
	}
}

// Env returns environment flags for other nodes
func (processes *Processes) Env() []string {
	var env []string
	for _, process := range processes.List {
		env = append(env, process.Info.Env()...)
	}
	return env
}

// Close closes all the processes and their resources
func (processes *Processes) Close() error {
	var errs []error
	for _, process := range processes.List {
		err := process.Close()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return utils.CombineErrors(errs...)
}

// ProcessInfo represents public information about the process
type ProcessInfo struct {
	Name      string
	ID        string
	Address   string
	Directory string
}

// Env returns process flags
func (info *ProcessInfo) Env() []string {
	name := strings.ToUpper(info.Name)

	var env []string
	if info.ID != "" {
		env = append(env, name+"_ID", info.ID)
	}
	if info.Address != "" {
		env = append(env, name+"_ADDR", info.Address)
	}
	if info.Directory != "" {
		env = append(env, name+"_DIR", info.Directory)
	}
	return env
}

// Process is a type for monitoring the process
type Process struct {
	Name       string
	Directory  string
	Executable string

	Info ProcessInfo

	Arguments map[string][]string

	Stdout io.Writer
	Stderr io.Writer

	stdout *os.File
	stderr *os.File
}

// NewProcess creates a process which can be run in the specified directory
func NewProcess(name, executable, directory string) (*Process, error) {
	stdout, err1 := os.OpenFile(filepath.Join(directory, "stderr.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	stderr, err2 := os.OpenFile(filepath.Join(directory, "stdout.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return &Process{
		Name:       name,
		Directory:  directory,
		Executable: executable,

		Arguments: map[string][]string{},

		Stdout: io.MultiWriter(os.Stdout, stdout),
		Stderr: io.MultiWriter(os.Stderr, stderr),

		stdout: stdout,
		stderr: stderr,
	}, utils.CombineErrors(err1, err2)
}

// Exec runs the process using the arguments for a given command
func (process *Process) Exec(ctx context.Context, command string) error {
	cmd := exec.CommandContext(ctx, process.Executable, process.Arguments[command]...)
	cmd.Dir = process.Directory
	cmd.Stdout, cmd.Stderr = process.Stdout, process.Stderr

	processgroup.Setup(cmd)

	if printCommands {
		fmt.Println("exec: ", strings.Join(cmd.Args, " "))
	}
	return cmd.Run()
}

// Close closes process resources
func (process *Process) Close() error {
	return utils.CombineErrors(
		process.stdout.Close(),
		process.stderr.Close(),
	)
}
