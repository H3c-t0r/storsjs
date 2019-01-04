#!/bin/bash
set -ueo pipefail

SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

go install -race -v storj.io/storj/cmd/{storj-sdk,satellite,storagenode,uplink,gateway}

# setup tmpdir for testfiles and cleanup
TMP=$(mktemp -d -t tmp.XXXXXXXXXX)
cleanup(){
	rm -rf "$TMP"
}
trap cleanup EXIT


# setup the network
storj-sdk -x --config-dir $TMP network setup

# run test-storj-sdk-aws.sh case
storj-sdk -x --config-dir $TMP network test bash $SCRIPTDIR/test-storj-sdk-aws.sh