// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"bytes"
	"os"
	"os/exec"
	"fmt"

	"github.com/zeebo/errs"
)

// UplinkError is class for minio errors
var UplinkError = errs.Class("uplink error")

// Uplink implements basic S3 Client with uplink
type Uplink struct {
	conf Config
}

// NewUplink creates new Client
func NewUplink(conf Config) (Client, error) {
	return &Uplink{conf}, nil
}

func (client *Uplink) cmd(subargs ...string) *exec.Cmd {
	args := []string{
		"--enc-key abc123",
	}

	args = append(args, subargs...)

	cmd := exec.Command("uplink", args...)
	return cmd
}

// MakeBucket makes a new bucket
func (client *Uplink) MakeBucket(bucket, location string) error {
	cmd := client.cmd("mb", "s3://"+bucket)
	_, err := cmd.Output()
	if err != nil {
		return UplinkError.Wrap(fullExitError(err))
	}
	return nil
}

// RemoveBucket removes a bucket
func (client *Uplink) RemoveBucket(bucket string) error {
	cmd := client.cmd("rb", "s3://"+bucket)
	_, err := cmd.Output()
	if err != nil {
		return UplinkError.Wrap(fullExitError(err))
	}
	return nil
}

// ListBuckets lists all buckets
func (client *Uplink) ListBuckets() ([]string, error) {
	cmd := client.cmd("ls")
	_, err := cmd.Output()
	if err != nil {
		return nil, UplinkError.Wrap(fullExitError(err))
	}

	// TODO: return something usefull
	return nil, nil
}

// Upload uploads object data to the specified path
func (client *Uplink) Upload(bucket, objectName string, data []byte) error {
	// TODO: add upload threshold
	cmd := client.cmd("put", "s3://"+bucket+"/"+objectName)
	cmd.Stdin = bytes.NewReader(data)
	_, err := cmd.Output()
	if err != nil {
		return UplinkError.Wrap(fullExitError(err))
	}
	return nil
}

// Download downloads object data
func (client *Uplink) Download(bucket, objectName string, buffer []byte) ([]byte, error) {
	cmd := client.cmd("cat", "s3://"+bucket+"/"+objectName)

	buf := &bufferWriter{buffer[:0]}
	cmd.Stdout = buf
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return nil, UplinkError.Wrap(fullExitError(err))
	}

	return buf.data, nil
}

// Delete deletes object
func (client *Uplink) Delete(bucket, objectName string) error {
	cmd := client.cmd("rm", "s3://"+bucket+"/"+objectName)
	_, err := cmd.Output()
	if err != nil {
		return UplinkError.Wrap(fullExitError(err))
	}
	return nil
}

// ListObjects lists objects
func (client *Uplink) ListObjects(bucket, prefix string) ([]string, error) {
	cmd := client.cmd("ls", "s3://"+bucket+"/"+prefix)
	data, err := cmd.Output()
	if err != nil {
		return nil, UplinkError.Wrap(fullExitError(err))
	}
	fmt.Printf("%v, %v\n", data, err)

	// TODO: return something usefull
	return nil, nil
}
