// Copyright (C) 2020 Storj Labs, Inc.
// See LICENSE for copying information.

package version

import _ "unsafe" // needed for go:linkname

//go:linkname buildTimestamp storj.io/private/version.buildTimestamp
var buildTimestamp string = "1704227581"

//go:linkname buildCommitHash storj.io/private/version.buildCommitHash
var buildCommitHash string = "7338249b9982eed9df6ac91ffdb88c34229764ec"

//go:linkname buildVersion storj.io/private/version.buildVersion
var buildVersion string = "v1.94.3"

//go:linkname buildRelease storj.io/private/version.buildRelease
var buildRelease string = "true"

// ensure that linter understands that the variables are being used.
func init() { use(buildTimestamp, buildCommitHash, buildVersion, buildRelease) }

func use(...interface{}) {}
