// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package main

// #include "uplink_definitions.h"
import "C"

import (
	libuplink "storj.io/storj/lib/uplink"
)

func main() {}

// Uplink is a scoped libuplink.Uplink.
type Uplink struct {
	scope
	lib *libuplink.Uplink
}

//export NewUplink
// NewUplink creates the uplink with the specified configuration and returns
// an error in cerr, when there is one.
//
// Caller must call CloseUplink to close associated resources.
func NewUplink(cfg C.UplinkConfig, cerr **C.char) C.Uplink {
	scope := rootScope("inmemory") // TODO: pass in as argument

	libcfg := &libuplink.Config{} // TODO: figure out a better name
	libcfg.Volatile.TLS.SkipPeerCAWhitelist = cfg.Volatile.TLS.SkipPeerCAWhitelist == 1

	lib, err := libuplink.NewUplink(scope.ctx, libcfg)
	if err != nil {
		*cerr = C.CString(err.Error())
		return C.Uplink{}
	}

	return C.Uplink{universe.Add(&Uplink{scope, lib})}
}

//export CloseUplink
// CloseUplink closes and frees the resources associated with uplink
func CloseUplink(uplinkHandle C.Uplink, cerr **C.char) {
	uplink, ok := universe.Get(uplinkHandle._handle).(*Uplink)
	if !ok {
		*cerr = C.CString("invalid uplink")
		return
	}
	universe.Del(uplinkHandle._handle)
	defer uplink.cancel()

	if err := uplink.lib.Close(); err != nil {
		*cerr = C.CString(err.Error())
		return
	}
}
