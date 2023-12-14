//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package datetimegroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"strings"
	"time"
)

const (
	dateTimeRFC1123JSON = `"` + time.RFC1123 + `"`
)

type dateTimeRFC1123 time.Time

func (t dateTimeRFC1123) MarshalJSON() ([]byte, error) {
	b := []byte(time.Time(t).Format(dateTimeRFC1123JSON))
	return b, nil
}

func (t dateTimeRFC1123) MarshalText() ([]byte, error) {
	b := []byte(time.Time(t).Format(time.RFC1123))
	return b, nil
}

func (t *dateTimeRFC1123) UnmarshalJSON(data []byte) error {
	p, err := time.Parse(dateTimeRFC1123JSON, strings.ToUpper(string(data)))
	*t = dateTimeRFC1123(p)
	return err
}

func (t *dateTimeRFC1123) UnmarshalText(data []byte) error {
	p, err := time.Parse(time.RFC1123, string(data))
	*t = dateTimeRFC1123(p)
	return err
}

func (t dateTimeRFC1123) String() string {
	return time.Time(t).Format(time.RFC1123)
}

func populateDateTimeRFC1123(m map[string]any, k string, t *time.Time) {
	if t == nil {
		return
	} else if azcore.IsNullValue(t) {
		m[k] = nil
		return
	} else if reflect.ValueOf(t).IsNil() {
		return
	}
	m[k] = (*dateTimeRFC1123)(t)
}

func unpopulateDateTimeRFC1123(data json.RawMessage, fn string, t **time.Time) error {
	if data == nil || strings.EqualFold(string(data), "null") {
		return nil
	}
	var aux dateTimeRFC1123
	if err := json.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	*t = (*time.Time)(&aux)
	return nil
}
