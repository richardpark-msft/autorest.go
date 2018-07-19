package arraygroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/satori/go.uuid"
)

// FooEnum enumerates the values for foo enum.
type FooEnum string

const (
	// Foo1 ...
	Foo1 FooEnum = "foo1"
	// Foo2 ...
	Foo2 FooEnum = "foo2"
	// Foo3 ...
	Foo3 FooEnum = "foo3"
)

// PossibleFooEnumValues returns an array of possible values for the FooEnum const type.
func PossibleFooEnumValues() []FooEnum {
	return []FooEnum{Foo1, Foo2, Foo3}
}

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ListBase64URL ...
type ListBase64URL struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
}

// ListBool ...
type ListBool struct {
	autorest.Response `json:"-"`
	Value             *[]bool `json:"value,omitempty"`
}

// ListByteArray ...
type ListByteArray struct {
	autorest.Response `json:"-"`
	Value             *[][]byte `json:"value,omitempty"`
}

// ListDate ...
type ListDate struct {
	autorest.Response `json:"-"`
	Value             *[]date.Date `json:"value,omitempty"`
}

// ListDateTime ...
type ListDateTime struct {
	autorest.Response `json:"-"`
	Value             *[]date.Time `json:"value,omitempty"`
}

// ListDateTimeRfc1123 ...
type ListDateTimeRfc1123 struct {
	autorest.Response `json:"-"`
	Value             *[]date.TimeRFC1123 `json:"value,omitempty"`
}

// ListFloat64 ...
type ListFloat64 struct {
	autorest.Response `json:"-"`
	Value             *[]float64 `json:"value,omitempty"`
}

// ListInt32 ...
type ListInt32 struct {
	autorest.Response `json:"-"`
	Value             *[]int32 `json:"value,omitempty"`
}

// ListInt64 ...
type ListInt64 struct {
	autorest.Response `json:"-"`
	Value             *[]int64 `json:"value,omitempty"`
}

// ListListString ...
type ListListString struct {
	autorest.Response `json:"-"`
	Value             *[][]string `json:"value,omitempty"`
}

// ListProduct ...
type ListProduct struct {
	autorest.Response `json:"-"`
	Value             *[]Product `json:"value,omitempty"`
}

// ListSetString ...
type ListSetString struct {
	autorest.Response `json:"-"`
	Value             *[]map[string]*string `json:"value,omitempty"`
}

// ListString ...
type ListString struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
}

// ListTimeSpan ...
type ListTimeSpan struct {
	autorest.Response `json:"-"`
	Value             *[]string `json:"value,omitempty"`
}

// ListUUID ...
type ListUUID struct {
	autorest.Response `json:"-"`
	Value             *[]uuid.UUID `json:"value,omitempty"`
}

// Product ...
type Product struct {
	Integer *int32  `json:"integer,omitempty"`
	String  *string `json:"string,omitempty"`
}
