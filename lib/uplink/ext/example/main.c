// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

#include <stdio.h>
#include <unistd.h>
#include "../uplink-cgo.h"

// gcc -o cgo-test-bin lib/uplink/ext/example/main.c lib/uplink/ext/uplink-cgo.so

int main() {
    char *err = "";

    struct Config uplinkConfig;
//    printf("getidversion\n");
//    struct IDVersion idVersion = {2};
    uplinkConfig.Volatile.IdentityVersion = GetIDVersion(0, &err);
    if (err != "") {
        printf("error: %s\n", err);
        return 1;
    }

//    printf("got idVersion %d\n", uplinkConfig.Volatile.IdentityVersion.Number);
    uplinkConfig.Volatile.TLS.SkipPeerCAWhitelist = true;

//    printf("newuplink\n");
    struct Uplink uplink = NewUplink(uplinkConfig, &err);

    if (err != "") {
        printf("error: %s\n", err);
        return 1;
    }


    printf("%d\n", uplink.Config.Volatile.IdentityVersion.Number);
    printf("%s\n", uplink.Config.Volatile.TLS.SkipPeerCAWhitelist ? "true" : "false");
    printf("%s\n", uplinkConfig.Volatile.TLS.SkipPeerCAWhitelist ? "true" : "false");
}