// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testplanet"
	"storj.io/storj/satellite/attribution"
)

func RunPlanet(t *testing.T, run func(ctx *testcontext.Context, planet *testplanet.Planet)) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	planet, err := testplanet.NewCustom(
		zaptest.NewLogger(t, zaptest.Level(zapcore.WarnLevel)),
		testplanet.Config{
			SatelliteCount:   1,
			StorageNodeCount: 6,
			UplinkCount:      1,
			Reconfigure:      testplanet.DisablePeerCAWhitelist,
		},
	)
	require.NoError(t, err)
	defer ctx.Check(planet.Shutdown)

	planet.Start(ctx)

	// make sure nodes are refreshed in db
	planet.Satellites[0].Discovery.Service.Refresh.TriggerWait()

	run(ctx, planet)
}

func TestC(t *testing.T) {
	ctx := testcontext.NewWithTimeout(t, 5*time.Minute)
	defer ctx.Cleanup()

	// TODO: compile libuplink only once
	libuplink := ctx.CompileShared(t, "uplink", "storj.io/storj/lib/uplinkc")

	currentdir, err := os.Getwd()
	require.NoError(t, err)

	definition := testcontext.Include{
		Header: filepath.Join(currentdir, "uplink_definitions.h"),
	}

	// TODO: combine all C tests into a single test executable
	//  (NB: a single failing test shouldn't fail the rest)
	ctests, err := filepath.Glob(filepath.Join("testdata", "*_test.c"))
	require.NoError(t, err)

	t.Run("ALL", func(t *testing.T) {
		for _, ctest := range ctests {
			ctest := ctest
			t.Run(filepath.Base(ctest), func(t *testing.T) {
				t.Parallel()

				testexe := ctx.CompileC(t, ctest, libuplink, definition)

				RunPlanet(t, func(ctx *testcontext.Context, planet *testplanet.Planet) {
					cmd := exec.Command(testexe)
					cmd.Dir = filepath.Dir(testexe)
					cmd.Env = append(os.Environ(),
						"SATELLITE_0_ADDR="+planet.Satellites[0].Addr(),
						"GATEWAY_0_API_KEY="+planet.Uplinks[0].APIKey[planet.Satellites[0].ID()],
					)

					out, err := cmd.CombinedOutput()
					if err != nil {
						t.Error(string(out))
						t.Fatal(err)
					} else {
						t.Log(string(out))
					}
				})
			})
		}
	})
}

func TestCBucketAttribution(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	newUUID := func() uuid.UUID {
		v, err := uuid.New()
		require.NoError(t, err)
		return *v
	}

	// TODO: compile libuplink only once
	libuplink := ctx.CompileShared(t, "uplink", "storj.io/storj/lib/uplinkc")

	currentdir, err := os.Getwd()
	require.NoError(t, err)

	definition := testcontext.Include{
		Header: filepath.Join(currentdir, "uplink_definitions.h"),
	}

	bucketTest := filepath.Join("testdata", "bucket_test.c")
	testexe := ctx.CompileC(t, bucketTest, libuplink, definition)

	RunPlanet(t, func(ctx *testcontext.Context, planet *testplanet.Planet) {
		checkPartnerID := func(expectedID *uuid.UUID) {
			projects, err := planet.Satellites[0].DB.Console().Projects().GetAll(ctx)
			require.NoError(t, err)
			require.Len(t, projects, 1)

			projectID := projects[0].ID

			// TODO: remove duplicated, hard-coded bucket name
			bucketName := []byte("test-bucket1")
			attrInfo, err := planet.Satellites[0].DB.Attribution().Get(ctx, projectID, bucketName)

			if expectedID != nil {
				require.NotNil(t, attrInfo)
				require.NoError(t, err)
				assert.Equal(t, attrInfo.PartnerID, *expectedID)
			} else {
				assert.True(t, attribution.ErrBucketNotAttributed.Has(err))
			}
		}

		{ // no partner id
			cmd := exec.Command(testexe)
			cmd.Dir = filepath.Dir(testexe)
			cmd.Env = append(os.Environ(),
				"SATELLITE_0_ADDR="+planet.Satellites[0].Addr(),
				"GATEWAY_0_API_KEY="+planet.Uplinks[0].APIKey[planet.Satellites[0].ID()],
			)

			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Error(string(out))
				t.Fatal(err)
			}
			t.Log(string(out))

			checkPartnerID(nil)
		}

		{ // valid partner id
			partnerID := newUUID()

			cmd := exec.Command(testexe)
			cmd.Dir = filepath.Dir(testexe)
			cmd.Env = append(os.Environ(),
				"SATELLITE_0_ADDR="+planet.Satellites[0].Addr(),
				"GATEWAY_0_API_KEY="+planet.Uplinks[0].APIKey[planet.Satellites[0].ID()],
				"PARTNER_ID_STR="+partnerID.String(),
			)

			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Error(string(out))
				t.Fatal(err)
			}
			t.Log(string(out))

			checkPartnerID(&partnerID)
		}
	})
}
