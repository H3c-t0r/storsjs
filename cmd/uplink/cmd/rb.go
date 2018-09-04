// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"storj.io/storj/pkg/cfgstruct"
	"storj.io/storj/pkg/process"
	"storj.io/storj/pkg/storage/meta"
	"storj.io/storj/pkg/utils"
	"storj.io/storj/storage"
)

var (
	rbCfg Config
	rbCmd = &cobra.Command{
		Use:   "rb",
		Short: "Remove an empty bucket",
		RunE:  deleteBucket,
	}
)

func init() {
	RootCmd.AddCommand(rbCmd)
	cfgstruct.Bind(rbCmd.Flags(), &rbCfg, cfgstruct.ConfDir(defaultConfDir))
	rbCmd.Flags().String("config", filepath.Join(defaultConfDir, "config.yaml"), "path to configuration")
}

func deleteBucket(cmd *cobra.Command, args []string) error {
	ctx := process.Ctx(cmd)

	identity, err := rbCfg.Load()
	if err != nil {
		return err
	}

	bs, err := rbCfg.GetBucketStore(ctx, identity)
	if err != nil {
		return err
	}

	if len(args) == 0 {
		fmt.Println("No bucket specified for deletion")
		return nil
	}

	u, err := utils.ParseURL(args[0])
	if err != nil {
		return err
	}

	_, err = bs.Get(ctx, u.Host)
	if err != nil {
		if storage.ErrKeyNotFound.Has(err) {
			fmt.Printf("Bucket not found: %s\n", u.Host)
			return nil
		}
		return err
	}

	o, err := bs.GetObjectStore(ctx, u.Host)
	if err != nil {
		return err
	}

	items, _, err := o.List(ctx, nil, nil, nil, true, 1, meta.None)
	if err != nil {
		return err
	}

	if len(items) > 0 {
		fmt.Printf("Bucket not empty: %s\n", u.Host)
		return nil
	}

	err = bs.Delete(ctx, u.Host)
	if err != nil {
		return err
	}

	fmt.Printf("Bucket deleted: %s\n", u.Host)

	return nil
}
