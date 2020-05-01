// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package paramgroupinggrouptest

import (
	"context"
	"generatortests/autorest/generated/paramgroupinggroup"
	"generatortests/helpers"
	"net/http"
	"testing"
)

func getParamGroupingOperations(t *testing.T) paramgroupinggroup.ParameterGroupingOperations {
	client, err := paramgroupinggroup.NewClient("http://localhost:3000", nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.ParameterGroupingOperations()
}

// PostMultiParamGroups - Post parameters from multiple different parameter groups
func TestPostMultiParamGroups(t *testing.T) {
	client := getParamGroupingOperations(t)
	result, err := client.PostMultiParamGroups(context.Background(), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostOptional - Post a bunch of optional parameters grouped
func TestPostOptional(t *testing.T) {
	client := getParamGroupingOperations(t)
	result, err := client.PostOptional(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostRequired - Post a bunch of required parameters grouped
func TestPostRequired(t *testing.T) {
	client := getParamGroupingOperations(t)
	result, err := client.PostRequired(context.Background(), paramgroupinggroup.ParameterGroupingPostRequiredParameters{
		Body:          1234,
		PathParameter: "path",
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}

// PostSharedParameterGroupObject - Post parameters with a shared parameter group object
func TestPostSharedParameterGroupObject(t *testing.T) {
	client := getParamGroupingOperations(t)
	result, err := client.PostSharedParameterGroupObject(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}