// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package fake

import (
	"net/http"
	"strings"
)

type nonRetriableError struct {
	error
}

func (nonRetriableError) NonRetriable() {
	// marker method
}

func contains[T comparable](s []T, v T) bool {
	for _, vv := range s {
		if vv == v {
			return true
		}
	}
	return false
}

func getHeaderValue(h http.Header, k string) string {
	v := h[k]
	if len(v) == 0 {
		return ""
	}
	return v[0]
}

func parseWithCast[T any](v string, parse func(v string) (T, error)) (T, error) {
	t, err := parse(v)
	if err != nil {
		return *new(T), err
	}
	return t, err
}

func splitHelper(s, sep string) []string {
	if s == "" {
		return nil
	}
	return strings.Split(s, sep)
}
