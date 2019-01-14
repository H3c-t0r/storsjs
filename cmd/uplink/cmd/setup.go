// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"storj.io/storj/internal/fpath"
	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/miniogw"
	"storj.io/storj/pkg/process"
	"storj.io/storj/pkg/provider"
)

var (
	uplinkSetupCmd = &cobra.Command{
		Use:         "setup",
		Short:       "Create an uplink config file",
		RunE:        cmdUplinkSetup,
		Annotations: map[string]string{"type": "setup"},
	}
	// gatewaySetupCmd = &cobra.Command{
	// 	Use:         "setup",
	// 	Short:       "Create an gateway config file",
	// 	RunE:        cmdGatewaySetup,
	// 	Annotations: map[string]string{"type": "setup"},
	// }
	uplinkCfg struct {
		CA            provider.CASetupConfig       `setup:"true"`
		Identity      provider.IdentitySetupConfig `setup:"true"`
		Overwrite     bool                         `default:"false" help:"whether to overwrite pre-existing configuration files" setup:"true"`
		SatelliteAddr string                       `default:"localhost:7778" help:"the address to use for the satellite" setup:"true"`

		Client miniogw.ClientConfig
		RS     miniogw.RSConfig
		Enc    miniogw.EncryptionConfig
	}

	cliConfDir *string
)

func init() {
	defaultUplinkConfDir := fpath.ApplicationDir("storj", "uplink")
	//defaultGatewayConfDir := fpath.ApplicationDir("storj", "gateway")

	dirParam := cfgstruct.FindConfigDirParam()
	if dirParam != "" {
		defaultUplinkConfDir = dirParam
		//defaultGatewayConfDir = dirParam
	}

	cliConfDir = CLICmd.PersistentFlags().String("config-dir", defaultUplinkConfDir, "main directory for setup configuration")
	//gwConfDir = GWCmd.PersistentFlags().String("config-dir", defaultGatewayConfDir, "main directory for setup configuration")

	CLICmd.AddCommand(uplinkSetupCmd)
	//GWCmd.AddCommand(gatewaySetupCmd)
	cfgstruct.BindSetup(uplinkSetupCmd.Flags(), &uplinkCfg, cfgstruct.ConfDir(defaultUplinkConfDir))
	//cfgstruct.BindSetup(gatewaySetupCmd.Flags(), &gatewayCfg, cfgstruct.ConfDir(defaultGatewayConfDir))
}

func cmdUplinkSetup(cmd *cobra.Command, args []string) (err error) {
	setupDir, err := filepath.Abs(*cliConfDir)
	if err != nil {
		return err
	}

	for _, flagname := range args {
		return fmt.Errorf("%s - Invalid flag. Pleas see --help", flagname)
	}

	valid, _ := fpath.IsValidSetupDir(setupDir)
	if !uplinkCfg.Overwrite && !valid {
		return fmt.Errorf("%s configuration already exists (%v). Rerun with --overwrite", "uplink", setupDir)
	}

	err = os.MkdirAll(setupDir, 0700)
	if err != nil {
		return err
	}

	defaultConfDir := fpath.ApplicationDir("storj", "uplink")
	// TODO: handle setting base path *and* identity file paths via args
	// NB: if base path is set this overrides identity and CA path options
	if setupDir != defaultConfDir {
		uplinkCfg.CA.CertPath = filepath.Join(setupDir, "ca.cert")
		uplinkCfg.CA.KeyPath = filepath.Join(setupDir, "ca.key")
		uplinkCfg.Identity.CertPath = filepath.Join(setupDir, "identity.cert")
		uplinkCfg.Identity.KeyPath = filepath.Join(setupDir, "identity.key")
	}
	err = provider.SetupIdentity(process.Ctx(cmd), uplinkCfg.CA, uplinkCfg.Identity)
	if err != nil {
		return err
	}

	o := map[string]interface{}{
		"identity.cert-path":     uplinkCfg.Identity.CertPath,
		"identity.key-path":      uplinkCfg.Identity.KeyPath,
		"client.pointer-db-addr": uplinkCfg.SatelliteAddr,
		"client.overlay-addr":    uplinkCfg.SatelliteAddr,
	}

	return process.SaveConfig(cmd.Flags(), filepath.Join(setupDir, "config.yaml"), o)
}
