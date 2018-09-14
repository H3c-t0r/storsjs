// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package cmd

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"gopkg.in/cheggaaa/pb.v1"

	"storj.io/storj/pkg/paths"
	"storj.io/storj/pkg/process"
	"storj.io/storj/pkg/storage/buckets"
	"storj.io/storj/pkg/storage/objects"
	"storj.io/storj/pkg/utils"
)

func init() {
	addCmd(&cobra.Command{
		Use:   "cp",
		Short: "Copies a local file or Storj object to another location locally or in Storj",
		RunE:  copyMain,
	})
}

func cleanAbsPath(p string) string {
	prefix := strings.HasSuffix(p, "/")
	p = path.Join("/", p)
	if !strings.HasSuffix(p, "/") && prefix {
		p += "/"
	}
	return p
}

// upload uploads args[0] from local machine to s3 compatible object args[1]
func upload(ctx context.Context, bs buckets.Store, srcFile string, destObj *url.URL) error {
	var err error
	if destObj.Scheme == "" {
		return fmt.Errorf("Invalid destination")
	}

	destObj.Path = cleanAbsPath(destObj.Path)
	// if object name not specified, default to filename
	if strings.HasSuffix(destObj.Path, "/") {
		destObj.Path = path.Join(destObj.Path, path.Base(srcFile))
	}

	var f *os.File
	if srcFile == "-" {
		f = os.Stdin
	} else {
		f, err = os.Open(srcFile)
		if err != nil {
			return err
		}
		defer utils.LogClose(f)
	}

	fi, err := f.Stat()
	if err != nil {
		return err
	}

	bar := pb.New(int(fi.Size())).SetUnits(pb.U_BYTES)
	bar.Start()

	reader := bar.NewProxyReader(f)
	defer utils.LogClose(reader)

	o, err := bs.GetObjectStore(ctx, destObj.Host)
	if err != nil {
		return err
	}

	meta := objects.SerializableMeta{}
	expTime := time.Time{}

	_, err = o.Put(ctx, paths.New(destObj.Path), reader, meta, expTime)
	if err != nil {
		return err
	}

	fmt.Printf("Created %s\n", destObj)

	return nil
}

// download downloads s3 compatible object args[0] to args[1] on local machine
func download(ctx context.Context, bs buckets.Store, srcObj *url.URL, destFile string) error {
	var err error
	if srcObj.Scheme == "" {
		return fmt.Errorf("Invalid source")
	}

	o, err := bs.GetObjectStore(ctx, srcObj.Host)
	if err != nil {
		return err
	}

	if fi, err := os.Stat(destFile); err == nil && fi.IsDir() {
		destFile = filepath.Join(destFile, filepath.Base(srcObj.Path))
	}

	var f *os.File
	if destFile == "-" {
		f = os.Stdout
	} else {
		f, err = os.Open(destFile)
		if err != nil {
			return err
		}
		defer utils.LogClose(f)
	}

	rr, _, err := o.Get(ctx, paths.New(srcObj.Path))
	if err != nil {
		return err
	}

	bar := pb.New(int(rr.Size())).SetUnits(pb.U_BYTES)
	bar.Start()

	r, err := rr.Range(ctx, 0, rr.Size())
	if err != nil {
		return err
	}
	defer utils.LogClose(r)

	reader := bar.NewProxyReader(r)
	defer utils.LogClose(reader)

	_, err = io.Copy(f, reader)
	if err != nil {
		return err
	}

	if destFile != "-" {
		fmt.Printf("Downloaded %s to %s\n", srcObj, destFile)
	}

	return nil
}

// copy copies s3 compatible object args[0] to s3 compatible object args[1]
func copy(ctx context.Context, bs buckets.Store, srcObj *url.URL, destObj *url.URL) error {
	o, err := bs.GetObjectStore(ctx, srcObj.Host)
	if err != nil {
		return err
	}

	rr, _, err := o.Get(ctx, paths.New(srcObj.Path))
	if err != nil {
		return err
	}

	bar := pb.New(int(rr.Size())).SetUnits(pb.U_BYTES)
	bar.Start()

	r, err := rr.Range(ctx, 0, rr.Size())
	if err != nil {
		return err
	}
	defer utils.LogClose(r)

	reader := bar.NewProxyReader(r)
	defer utils.LogClose(reader)

	if destObj.Host != srcObj.Host {
		o, err = bs.GetObjectStore(ctx, destObj.Host)
		if err != nil {
			return err
		}
	}

	meta := objects.SerializableMeta{}
	expTime := time.Time{}

	destObj.Path = cleanAbsPath(destObj.Path)
	// if destination object name not specified, default to source object name
	if strings.HasSuffix(destObj.Path, "/") {
		destObj.Path = path.Join(destObj.Path, path.Base(srcObj.Path))
	}

	_, err = o.Put(ctx, paths.New(destObj.Path), reader, meta, expTime)
	if err != nil {
		return err
	}

	fmt.Printf("%s copied to %s\n", srcObj, destObj)

	return nil
}

// copyMain is the function executed when cpCmd is called
func copyMain(cmd *cobra.Command, args []string) (err error) {
	if len(args) == 0 {
		return fmt.Errorf("No object specified for copy")
	}
	if len(args) == 1 {
		return fmt.Errorf("No destination specified")
	}

	ctx := process.Ctx(cmd)

	u0, err := utils.ParseURL(args[0])
	if err != nil {
		return err
	}

	u1, err := utils.ParseURL(args[1])
	if err != nil {
		return err
	}

	bs, err := cfg.BucketStore(ctx)
	if err != nil {
		return err
	}

	// if uploading
	if u0.Scheme == "" {
		if u1.Host == "" {
			return fmt.Errorf("No bucket specified. Please use format sj://bucket/")
		}

		return upload(ctx, bs, args[0], u1)
	}

	// if downloading
	if u1.Scheme == "" {
		if u0.Host == "" {
			return fmt.Errorf("No bucket specified. Please use format sj://bucket/")
		}

		return download(ctx, bs, u0, args[1])
	}

	// if copying from one remote location to another
	return copy(ctx, bs, u0, u1)
}
