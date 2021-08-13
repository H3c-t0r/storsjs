// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package ultest

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"storj.io/storj/cmd/uplinkng/ulloc"
)

// Result captures all the output of running a command for inspection.
type Result struct {
	Stdout string
	Stderr string
	Ok     bool
	Err    error
	Files  []File
}

// RequireSuccess fails if the Result did not observe a successful execution.
func (r Result) RequireSuccess(t *testing.T) {
	if !r.Ok {
		errs := parseErrors(r.Stdout)
		require.FailNow(t, "test did not run successfully:",
			"%s", strings.Join(errs, "\n"))
	}
	require.NoError(t, r.Err)
}

// RequireFailure fails if the Result did not observe a failed execution.
func (r Result) RequireFailure(t *testing.T) Result {
	require.False(t, r.Ok && r.Err == nil, "command ran with no error")
	return r
}

// RequireStdout requires that the execution wrote to stdout the provided string.
// Blank lines are ignored and all lines are space trimmed for the comparison.
func (r Result) RequireStdout(t *testing.T, stdout string) Result {
	require.Equal(t, trimNewlineSpaces(stdout), trimNewlineSpaces(r.Stdout))
	return r
}

// RequireStderr requires that the execution wrote to stderr the provided string.
// Blank lines are ignored and all lines are space trimmed for the comparison.
func (r Result) RequireStderr(t *testing.T, stderr string) Result {
	require.Equal(t, trimNewlineSpaces(stderr), trimNewlineSpaces(r.Stderr))
	return r
}

// RequireFiles requires that the set of files provided are all of the files that
// existed at the end of the execution. It assumes any passed in files with no
// contents contain the filename as the contents instead.
func (r Result) RequireFiles(t *testing.T, files ...File) Result {
	require.Equal(t, canonicalizeFiles(files), r.Files)
	return r
}

// RequireLocalFiles requires that the set of files provided are all of the
// local files that existed at the end of the execution. It assumes any passed
// in files with no contents contain the filename as the contents instead.
func (r Result) RequireLocalFiles(t *testing.T, files ...File) Result {
	require.Equal(t, canonicalizeFiles(files), filterFiles(r.Files, fileIsLocal))
	return r
}

// RequireRemoteFiles requires that the set of files provided are all of the
// remote files that existed at the end of the execution. It assumes any passed
// in files with no contents contain the filename as the contents instead.
func (r Result) RequireRemoteFiles(t *testing.T, files ...File) Result {
	require.Equal(t, canonicalizeFiles(files), filterFiles(r.Files, fileIsRemote))
	return r
}

func filterFiles(files []File, match func(File) bool) (out []File) {
	for _, file := range files {
		if match(file) {
			out = append(out, file)
		}
	}
	return out
}

func canonicalizeFiles(files []File) (out []File) {
	out = append(out, files...)
	sort.Slice(out, func(i, j int) bool { return out[i].less(out[j]) })

	for i := range out {
		if out[i].Contents == "" {
			out[i].Contents = out[i].Loc
		}
	}

	return out
}

func fileIsLocal(file File) bool {
	loc, _ := ulloc.Parse(file.Loc)
	return loc.Local()
}

func fileIsRemote(file File) bool {
	loc, _ := ulloc.Parse(file.Loc)
	return loc.Remote()
}

func parseErrors(s string) []string {
	lines := strings.Split(s, "\n")
	start := 0
	for i, line := range lines {
		if line == "Errors:" {
			start = i + 1
		} else if len(line) > 0 && line[0] != ' ' {
			return lines[start:i]
		}
	}
	return nil
}

func trimNewlineSpaces(s string) string {
	lines := strings.Split(s, "\n")

	j := 0
	for _, line := range lines {
		if trimmed := strings.TrimSpace(line); len(trimmed) > 0 {
			lines[j] = trimmed
			j++
		}
	}
	return strings.Join(lines[:j], "\n")
}

// File represents a file existing either locally or remotely.
type File struct {
	Loc      string
	Contents string
}

func (f File) less(g File) bool {
	fl, _ := ulloc.Parse(f.Loc)
	gl, _ := ulloc.Parse(g.Loc)
	return fl.Less(gl)
}
