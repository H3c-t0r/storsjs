// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

#include <stdio.h>
#include <unistd.h>
#include <string.h>
#include <time.h>
#include "unit_test.h"
#include "../../uplink-cgo.h"
#include "helpers.h"

void TestBucket(void)
{
    char *_err = "";
    char **err = &_err;
    char *bucket_name = "TestBucket";

    // Open Project
    ProjectRef_t ref_project = OpenTestProject(err);
    TEST_ASSERT_EQUAL_STRING("", *err);

    // TODO: test with different bucket configs
    CreateBucket(ref_project, bucket_name, NULL, err);
    TEST_ASSERT_EQUAL_STRING("", *err);

    BucketRef_t ref_bucket = OpenBucket(ref_project, bucket_name, NULL, err);
    TEST_ASSERT_EQUAL_STRING("", *err);


    char *object_paths[] = {"TestObject1","TestObject2","TestObject3","TestObject4"};
    int num_of_objects = 4;

    // Create objects
    char *str_data = "testing data 123";
    for (int i=0; i < num_of_objects; i++) {
        Object_t *object = malloc(sizeof(Object_t));
        Bytes_t *data = BytesFromString(str_data);

        create_test_object(ref_bucket, object_paths[i], object, data, err);
        TEST_ASSERT_EQUAL_STRING("", *err);
        free(object);
        free(data);
    }

    // List objects
    // TODO: test list options
    ObjectList_t objects_list = ListObjects(ref_bucket, NULL, err);
    TEST_ASSERT_EQUAL_STRING("", *err);
    // TODO: add assertions

    // TODO: add assertions for metadata

    // TODO: Open Object
}

int main(int argc, char *argv[])
{
    TestBucket();
}
