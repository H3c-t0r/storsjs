// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package version

import _ "unsafe" // needed for go:linkname

//go:linkname buildTimestamp storj.io/private/version.buildTimestamp
var buildTimestamp string = "1651134574"

//go:linkname buildCommitHash storj.io/private/version.buildCommitHash
var buildCommitHash string = "04b22f24eb3e1f01c0e7b323d6333d4e1c775bbd"

//go:linkname buildVersion storj.io/private/version.buildVersion
var buildVersion string = "v1.54.0-rc"

//go:linkname buildRelease storj.io/private/version.buildRelease
var buildRelease string = "true"

// ensure that linter understands that the variables are being used.
func init() { use(buildTimestamp, buildCommitHash, buildVersion, buildRelease) }

func use(...interface{}) {}
