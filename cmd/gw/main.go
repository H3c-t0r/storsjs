// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/miniogw"
	"storj.io/storj/pkg/process"
	"storj.io/storj/pkg/provider"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gw",
		Short: "Gateway",
	}
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run the gateway",
		RunE:  cmdRun,
	}
	setupCmd = &cobra.Command{
		Use:   "setup",
		Short: "Create config files",
		RunE:  cmdSetup,
	}

	runCfg   miniogw.Config
	setupCfg struct {
		BasePath    string `default:"$CONFDIR" help:"base path for setup"`
		Concurrency uint   `default:"4" help:"number of concurrent workers for certificate authority generation"`
		CA          provider.CAConfig
		Identity    provider.IdentityConfig
		Overwrite bool   `default:"false" help:"whether to overwrite pre-existing configuration files"`

	}

	defaultConfDir = "$HOME/.storj/gw"
)

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(setupCmd)
	cfgstruct.Bind(runCmd.Flags(), &runCfg, cfgstruct.ConfDir(defaultConfDir))
	cfgstruct.Bind(setupCmd.Flags(), &setupCfg, cfgstruct.ConfDir(defaultConfDir))
}

func cmdRun(cmd *cobra.Command, args []string) (err error) {
	return runCfg.Run(process.Ctx(cmd))
}

func cmdSetup(cmd *cobra.Command, args []string) (err error) {
	_, err = os.Stat(setupCfg.BasePath)
	if !setupCfg.Overwrite && err == nil {
		fmt.Println("A gw configuration already exists. Rerun with --overwrite")
		return nil
	}

	err = os.MkdirAll(setupCfg.BasePath, 0700)
	if err != nil {
		return err
	}

	// Load or create a certificate authority
	ca, err := setupCfg.CA.LoadOrCreate(nil, 4)
	if err != nil {
		return err
	}
	// Load or create identity from CA
	_, err = setupCfg.Identity.LoadOrCreate(ca)
	if err != nil {
		return err
	}

	return process.SaveConfig(runCmd.Flags(),
		filepath.Join(setupCfg.BasePath, "config.yaml"), nil)
}

func main() {
	runCmd.Flags().String("config",
		filepath.Join(defaultConfDir, "config.yaml"), "path to configuration")
	process.Exec(rootCmd)
}
