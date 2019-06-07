// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #ifndef STORJ_HEADERS
//   #define STORJ_HEADERS
//   #include "c/headers/main.h"
// #endif
import "C"
import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/lib/uplink"
)

// C types
type Cchar = *C.char

// Ref types
type CAPIKeyRef = C.APIKeyRef_t

var (
	cLibDir, cSrcDir, cTestsDir, libuplink string

	testConfig = new(uplink.Config)
)

func init() {
	// TODO: is there a cleaner way to do this?
	_, thisFile, _, _ := runtime.Caller(0)
	cLibDir = filepath.Join(filepath.Dir(thisFile), "c")
	cSrcDir = filepath.Join(cLibDir, "src")
	cTestsDir = filepath.Join(cLibDir, "tests")
	libuplink = filepath.Join(cLibDir, "..", "uplink-cgo.so")

	testConfig.Volatile.TLS.SkipPeerCAWhitelist = true
}

func runCTests(t *testing.T, ctx *testcontext.Context, envVars []string, srcGlobs ...string) {
	srcGlobs = append([]string{
		libuplink,
		filepath.Join(cTestsDir, "unity.c"),
		filepath.Join(cTestsDir, "helpers.c"),
		filepath.Join(cSrcDir, "*.c"),
	}, srcGlobs...)
	testBinPath := ctx.CompileC(srcGlobs...)
	commandPath := testBinPath

	if dir, ok := os.LookupEnv("STORJ_DEBUG"); ok {
		err := copyFile(testBinPath, filepath.Join(dir, t.Name()))
		require.NoError(t, err)
	}

	cmd := exec.Command(commandPath)
	cmd.Env = append(os.Environ(), envVars...)

	out, err := cmd.CombinedOutput()
	t.Log(string(out))
	require.NoError(t, err)
}

func copyFile(src, dest string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dest, input, 0755)
	if err != nil {
		return err
	}
	return nil
}

func runCTest(t *testing.T, ctx *testcontext.Context, filename string, envVars ...string) {
	runCTests(t, ctx, envVars, filepath.Join(cLibDir, "tests", filename))
}

func startTestPlanet(t *testing.T, ctx *testcontext.Context) *testplanet.Planet {
	planet, err := testplanet.NewCustom(
		zap.NewNop(),
		testplanet.Config{
			SatelliteCount:     1,
			StorageNodeCount:   8,
			UplinkCount:        0,
			UsePeerCAWhitelist: false,
		},
	)
	require.NoError(t, err)

	planet.Start(ctx)
	return planet
}

func stringToCCharPtr(str string) *C.char {
	return (*C.char)(unsafe.Pointer(C.CString(str)))
}

func cCharToGoString(cchar *C.char) string {
	return C.GoString(cchar)
}

