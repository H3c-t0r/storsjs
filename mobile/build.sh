#!/bin/bash

# This script will build libuplink-android.aar library from scratch
# Required:
# * ANDROID_HOME set with NDK available
# * go

if [ -z "$ANDROID_HOME" ]
then
      echo "\$ANDROID_HOME is not set"
      exit 1
fi

# setup tmpdir for testfiles and cleanup
TMP=$(mktemp -d -t tmp.XXXXXXXXXX)
cleanup(){
	rm -rf "$TMP"
}
trap cleanup EXIT

OUTPUT=$PWD

go mod vendor

# go knows where our gopath is
export GOPATH=$TMP

# set go modules to default behavior
export GO111MODULE=off

mkdir "$GOPATH/src"

mv $PWD/vendor/* "$GOPATH/src"

# symlink doesn't look to be working with gomobile
# ln -s "$PWD/../storj" "$GOPATH/src/storj.io/storj"
rsync -am --stats --exclude=".*" "$PWD/../storj" "$GOPATH/src/storj.io/"

cd "$GOPATH/src/storj.io/storj"

go get golang.org/x/mobile/cmd/gomobile

$GOPATH/bin/gomobile init

$GOPATH/bin/gomobile bind -v -target android -o "$OUTPUT/libuplink-android.aar" -javapkg io.storj.libuplink storj.io/storj/mobile