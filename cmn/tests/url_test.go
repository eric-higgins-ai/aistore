// Package test provides tests for common low-level types and utilities for all aistore projects
/*
 * Copyright (c) 2018-2026, NVIDIA CORPORATION. All rights reserved.
 */
package tests_test

import (
	"testing"

	"github.com/NVIDIA/aistore/cmn"
	"github.com/NVIDIA/aistore/tools/tassert"
)

func TestParseURLScheme(t *testing.T) {
	testCases := []struct{ expectedScheme, expectedAddress, url string }{
		{"http", "localhost:8080", "http://localhost:8080"},
		{"https", "localhost", "https://localhost"},
		{"", "localhost:8080", "localhost:8080"},
	}

	for _, tc := range testCases {
		scheme, address := cmn.ParseURLScheme(tc.url)
		tassert.Errorf(t, scheme == tc.expectedScheme, "expected scheme %s, got %s", tc.expectedScheme, scheme)
		tassert.Errorf(t, address == tc.expectedAddress, "expected address %s, got %s", tc.expectedAddress, address)
	}
}
