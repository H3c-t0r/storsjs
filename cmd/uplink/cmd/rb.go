// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"storj.io/common/fpath"
)

func init() {
	addCmd(&cobra.Command{
		Use:   "rb sj://BUCKET",
		Short: "Remove an empty bucket",
		RunE:  deleteBucket,
		Args:  cobra.ExactArgs(1),
	}, RootCmd)
}

func deleteBucket(cmd *cobra.Command, args []string) error {
	ctx, _ := withTelemetry(cmd)

	if len(args) == 0 {
		return fmt.Errorf("no bucket specified for deletion")
	}

	dst, err := fpath.New(args[0])
	if err != nil {
		return err
	}

	if dst.IsLocal() {
		return fmt.Errorf("no bucket specified, use format sj://bucket/")
	}

	if dst.Path() != "" {
		return fmt.Errorf("nested buckets not supported, use format sj://bucket/")
	}

	project, err := cfg.getProject(ctx, false)
	if err != nil {
		return convertError(err, dst)
	}
	defer closeProject(project)

	if _, err := project.DeleteBucket(ctx, dst.Bucket()); err != nil {
		return convertError(err, dst)
	}

	fmt.Printf("Bucket %s deleted\n", dst.Bucket())

	return nil
}
