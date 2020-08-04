// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgrouptest

import (
	"context"
	"generatortests/autorest/generated/nonstringenumgroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getIntOperations(t *testing.T) nonstringenumgroup.IntOperations {
	client, err := nonstringenumgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatal(err)
	}
	return client.IntOperations()
}

// Get - Get an int enum
func TestIntGet(t *testing.T) {
	client := getIntOperations(t)
	result, err := client.Get(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	helpers.DeepEqualOrFatal(t, result.Value, nonstringenumgroup.IntEnumFourHundredTwentyNine.ToPtr())
}

// Put - Put an int enum
func TestIntPut(t *testing.T) {
	client := getIntOperations(t)
	result, err := client.Put(context.Background(), &nonstringenumgroup.IntPutOptions{
		Input: nonstringenumgroup.IntEnumTwoHundred.ToPtr(),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result.RawResponse, http.StatusOK)
}