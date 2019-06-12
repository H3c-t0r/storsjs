// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

#include <string.h>
#include <stdlib.h>

#include "require.h"
#include "uplink.h"
#include "helpers2.h"

void handle_project(Project project);

int main(int argc, char *argv[]) {
    with_test_project(&handle_project);
}

void handle_project(Project project) {
    char *_err = "";
    char **err = &_err;

    char *bucket_name = "TestBucket";

    BucketConfig config = test_bucket_config();
    BucketInfo info = create_bucket(project, bucket_name, &config, err);
    require_noerror(*err);
    free_bucket_info(&info);

    EncryptionAccess access = {};
    memcpy(&access.key[0], "hello", 5);
    Bucket bucket = open_bucket(project, bucket_name, access, err);
    require_noerror(*err);
    {
        char *object_paths[] = {"TestObject1","TestObject2","TestObject3","TestObject4"};
        int num_of_objects = 4;

        char *data = "testing data 123";
        //for(int i = 0; i < num_of_objects; i++) {
        //    
        //}
    }
    close_bucket(bucket, err);
    require_noerror(*err);
}
