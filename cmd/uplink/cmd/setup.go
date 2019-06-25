// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeebo/errs"
	"go.uber.org/zap"

	"storj.io/storj/internal/fpath"
	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/process"
	"storj.io/storj/uplink"
)

var (
	setupCmd = &cobra.Command{
		Use:         "setup",
		Short:       "Create an uplink config file",
		RunE:        cmdSetup,
		Annotations: map[string]string{"type": "setup"},
	}

	setupCfg UplinkFlags
	confDir  string
	defaults cfgstruct.BindOpt

	// Error is the default uplink setup errs class
	Error = errs.Class("uplink setup error")
)

func init() {
	defaultConfDir := fpath.ApplicationDir("storj", "uplink")
	cfgstruct.SetupFlag(zap.L(), RootCmd, &confDir, "config-dir", defaultConfDir, "main directory for uplink configuration")
	defaults = cfgstruct.DefaultsFlag(RootCmd)
	RootCmd.AddCommand(setupCmd)
	process.Bind(setupCmd, &setupCfg, defaults, cfgstruct.ConfDir(confDir), cfgstruct.SetupMode())
}

func cmdSetup(cmd *cobra.Command, args []string) (err error) {
	// Ensure use the default port if the user only specifies a host.
	err = ApplyDefaultHostAndPortToAddrFlag(cmd, "satellite-addr")
	if err != nil {
		return err
	}

	setupDir, err := filepath.Abs(confDir)
	if err != nil {
		return err
	}

	valid, _ := fpath.IsValidSetupDir(setupDir)
	if !valid {
		return fmt.Errorf("uplink configuration already exists (%v)", setupDir)
	}

	err = os.MkdirAll(setupDir, 0700)
	if err != nil {
		return err
	}

	// override is required because the default value of Enc.KeyFilepath is ""
	// and setting the value directly in setupCfg.Enc.KeyFiletpathon will set the
	// value in the config file but commented out.
	usedEncryptionKeyFilepath := setupCfg.Enc.KeyFilepath
	if usedEncryptionKeyFilepath == "" {
		usedEncryptionKeyFilepath = filepath.Join(setupDir, ".encryption.key")
	}

	if setupCfg.NonInteractive {
		return cmdSetupNonInteractive(cmd, setupDir, usedEncryptionKeyFilepath)
	}

	return cmdSetupInteractive(cmd, setupDir, usedEncryptionKeyFilepath)
}

// cmdSetupNonInteractive sets up uplink non-interactively.
//
// encryptionKeyFilepath should be set to the filepath indicated by the user or
// or to a default path whose directory tree exists.
func cmdSetupNonInteractive(cmd *cobra.Command, setupDir string, encryptionKeyFilepath string) error {
	if setupCfg.Enc.EncryptionKey != "" {
		err := uplink.SaveEncryptionKey(setupCfg.Enc.EncryptionKey, encryptionKeyFilepath)
		if err != nil {
			return err
		}
	}

	override := map[string]interface{}{
		"enc.key-filepath": encryptionKeyFilepath,
	}

	err := process.SaveConfigWithAllDefaults(
		cmd.Flags(), filepath.Join(setupDir, process.DefaultCfgFilename), override)
	if err != nil {
		return err
	}

	if setupCfg.Enc.EncryptionKey != "" {
		_, _ = fmt.Printf("Your encryption key is saved to: %s\n", encryptionKeyFilepath)
	}

	return nil
}

// cmdSetupInteractive sets up uplink interactively.
//
// encryptionKeyFilepath should be set to the filepath indicated by the user or
// or to a default path whose directory tree exists.
func cmdSetupInteractive(cmd *cobra.Command, setupDir string, encryptionKeyFilepath string) error {
	_, err := fmt.Print(`
Pick satellite to use:
	[1] us-central-1.tardigrade.io
	[2] europe-west-1.tardigrade.io
	[3] asia-east-1.tardigrade.io
Please enter numeric choice or enter satellite address manually [1]: `)
	if err != nil {
		return err
	}
	satellites := []string{"us-central-1.tardigrade.io", "europe-west-1.tardigrade.io", "asia-east-1.tardigrade.io"}
	var satelliteAddress string
	n, err := fmt.Scanln(&satelliteAddress)
	if err != nil {
		if n == 0 {
			// fmt.Scanln cannot handle empty input
			satelliteAddress = satellites[0]
		} else {
			return err
		}
	}

	// TODO add better validation
	if satelliteAddress == "" {
		return errs.New("satellite address cannot be empty")
	} else if len(satelliteAddress) == 1 {
		switch satelliteAddress {
		case "1":
			satelliteAddress = satellites[0]
		case "2":
			satelliteAddress = satellites[1]
		case "3":
			satelliteAddress = satellites[2]
		default:
			return errs.New("Satellite address cannot be one character")
		}
	}

	satelliteAddress, err = ApplyDefaultHostAndPortToAddr(
		satelliteAddress, cmd.Flags().Lookup("satellite-addr").Value.String())
	if err != nil {
		return err
	}

	_, err = fmt.Print("Enter your API key: ")
	if err != nil {
		return err
	}
	var apiKey string
	n, err = fmt.Scanln(&apiKey)
	if err != nil && n != 0 {
		return err
	}

	if apiKey == "" {
		return errs.New("API key cannot be empty")
	}

	humanReadableKey, err := cfgstruct.PromptForEncryptionKey()
	if err != nil {
		return err
	}

	err = uplink.SaveEncryptionKey(humanReadableKey, encryptionKeyFilepath)
	if err != nil {
		return err
	}

	var override = map[string]interface{}{
		"api-key":          apiKey,
		"satellite-addr":   satelliteAddress,
		"enc.key-filepath": encryptionKeyFilepath,
	}

	err = process.SaveConfigWithAllDefaults(
		cmd.Flags(), filepath.Join(setupDir, process.DefaultCfgFilename), override)
	if err != nil {
		return nil
	}

	// if there is an error with this we cannot do that much and the setup process
	// has ended OK, so we ignore it.
	_, _ = fmt.Printf(`
Your encryption key is saved to: %s

Your Uplink CLI is configured and ready to use!

Some things to try next:

* Run 'uplink --help' to see the operations that can be performed

* See https://github.com/storj/docs/blob/master/Uplink-CLI.md#usage for some example commands
	`, encryptionKeyFilepath)

	return nil
}

// ApplyDefaultHostAndPortToAddrFlag applies the default host and/or port if either is missing in the specified flag name.
func ApplyDefaultHostAndPortToAddrFlag(cmd *cobra.Command, flagName string) error {
	flag := cmd.Flags().Lookup(flagName)
	if flag == nil {
		// No flag found for us to handle.
		return nil
	}

	address, err := ApplyDefaultHostAndPortToAddr(flag.Value.String(), flag.DefValue)
	if err != nil {
		return Error.Wrap(err)
	}

	if flag.Value.String() == address {
		// Don't trip the flag set bit
		return nil
	}

	return Error.Wrap(flag.Value.Set(address))
}

// ApplyDefaultHostAndPortToAddr applies the default host and/or port if either is missing in the specified address.
func ApplyDefaultHostAndPortToAddr(address, defaultAddress string) (string, error) {
	defaultHost, defaultPort, err := net.SplitHostPort(defaultAddress)
	if err != nil {
		return "", Error.Wrap(err)
	}

	addressParts := strings.Split(address, ":")
	numberOfParts := len(addressParts)

	if numberOfParts > 1 && len(addressParts[0]) > 0 && len(addressParts[1]) > 0 {
		// address is host:port so skip applying any defaults.
		return address, nil
	}

	// We are missing a host:port part. Figure out which part we are missing.
	indexOfPortSeparator := strings.Index(address, ":")
	lengthOfFirstPart := len(addressParts[0])

	if indexOfPortSeparator < 0 {
		if lengthOfFirstPart == 0 {
			// address is blank.
			return defaultAddress, nil
		}
		// address is host
		return net.JoinHostPort(addressParts[0], defaultPort), nil
	}

	if indexOfPortSeparator == 0 {
		// address is :1234
		return net.JoinHostPort(defaultHost, addressParts[1]), nil
	}

	// address is host:
	return net.JoinHostPort(addressParts[0], defaultPort), nil
}
