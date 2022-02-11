// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package nonstringenumgroup

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func newFloatClient() *FloatClient {
	return NewFloatClient(nil)
}

// Get - Get a float enum
func TestFloatGet(t *testing.T) {
	client := newFloatClient()
	result, err := client.Get(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := cmp.Diff(result.Value, FloatEnumFourHundredTwentyNine1.ToPtr()); r != "" {
		t.Fatal(r)
	}
}

// Put - Put a float enum
func TestFloatPut(t *testing.T) {
	client := newFloatClient()
	result, err := client.Put(context.Background(), &FloatClientPutOptions{
		Input: FloatEnumTwoHundred4.ToPtr(),
	})
	if err != nil {
		t.Fatal(err)
	}
	if *result.Value != "Nice job posting a float enum" {
		t.Fatalf("unexpected value %s", *result.Value)
	}
}
