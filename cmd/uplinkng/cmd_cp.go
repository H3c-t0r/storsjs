// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"fmt"
	"io"
	"strconv"

	progressbar "github.com/cheggaaa/pb/v3"
	"github.com/zeebo/clingy"
	"github.com/zeebo/errs"

	"storj.io/storj/cmd/uplinkng/ulext"
	"storj.io/storj/cmd/uplinkng/ulfs"
	"storj.io/storj/cmd/uplinkng/ulloc"
)

type cmdCp struct {
	ex ulext.External

	access    string
	recursive bool
	dryrun    bool
	progress  bool

	source ulloc.Location
	dest   ulloc.Location
}

func newCmdCp(ex ulext.External) *cmdCp {
	return &cmdCp{ex: ex}
}

func (c *cmdCp) Setup(params clingy.Parameters) {
	c.access = params.Flag("access", "Which access to use", "").(string)
	c.recursive = params.Flag("recursive", "Peform a recursive copy", false,
		clingy.Short('r'),
		clingy.Transform(strconv.ParseBool),
	).(bool)
	c.dryrun = params.Flag("dryrun", "Print what operations would happen but don't execute them", false,
		clingy.Transform(strconv.ParseBool),
	).(bool)
	c.progress = params.Flag("progress", "Show a progress bar when possible", true,
		clingy.Transform(strconv.ParseBool),
	).(bool)

	c.source = params.Arg("source", "Source to copy", clingy.Transform(ulloc.Parse)).(ulloc.Location)
	c.dest = params.Arg("dest", "Desination to copy", clingy.Transform(ulloc.Parse)).(ulloc.Location)
}

func (c *cmdCp) Execute(ctx clingy.Context) error {
	fs, err := c.ex.OpenFilesystem(ctx, c.access)
	if err != nil {
		return err
	}
	defer func() { _ = fs.Close() }()

	if c.recursive {
		return c.copyRecursive(ctx, fs)
	}
	return c.copyFile(ctx, fs, c.source, c.dest, c.progress)
}

func (c *cmdCp) copyRecursive(ctx clingy.Context, fs ulfs.Filesystem) error {
	if c.source.Std() || c.dest.Std() {
		return errs.New("cannot recursively copy to stdin/stdout")
	}

	iter, err := fs.ListObjects(ctx, c.source, true)
	if err != nil {
		return err
	}

	anyFailed := false
	for iter.Next() {
		rel, err := c.source.RelativeTo(iter.Item().Loc)
		if err != nil {
			return err
		}

		source := iter.Item().Loc
		dest := c.dest.AppendKey(rel)

		if err := c.copyFile(ctx, fs, source, dest, false); err != nil {
			fmt.Fprintln(ctx.Stderr(), copyVerb(source, dest), "failed:", err.Error())
			anyFailed = true
		}
	}

	if err := iter.Err(); err != nil {
		return errs.Wrap(err)
	} else if anyFailed {
		return errs.New("some downloads failed")
	}
	return nil
}

func (c *cmdCp) copyFile(ctx clingy.Context, fs ulfs.Filesystem, source, dest ulloc.Location, progress bool) error {
	if isDir := fs.IsLocalDir(ctx, dest); isDir {
		base, ok := source.Base()
		if !ok {
			return errs.New("destination is a directory and cannot find base name for %q", source)
		}
		dest = dest.AppendKey(base)
	}

	if !source.Std() && !dest.Std() {
		fmt.Fprintln(ctx.Stdout(), copyVerb(source, dest), source, "to", dest)
	}

	if c.dryrun {
		return nil
	}

	rh, err := fs.Open(ctx, source)
	if err != nil {
		return err
	}
	defer func() { _ = rh.Close() }()

	wh, err := fs.Create(ctx, dest)
	if err != nil {
		return err
	}
	defer func() { _ = wh.Abort() }()

	var bar *progressbar.ProgressBar
	var writer io.Writer = wh

	if length := rh.Info().ContentLength; progress && length >= 0 && !c.dest.Std() {
		bar = progressbar.New64(length).SetWriter(ctx.Stdout())
		writer = bar.NewProxyWriter(writer)
		bar.Start()
		defer bar.Finish()
	}

	if _, err := io.Copy(writer, rh); err != nil {
		return errs.Combine(err, wh.Abort())
	}
	return errs.Wrap(wh.Commit())
}

func copyVerb(source, dest ulloc.Location) string {
	switch {
	case dest.Remote():
		return "upload"
	case source.Remote():
		return "download"
	default:
		return "copy"
	}
}
