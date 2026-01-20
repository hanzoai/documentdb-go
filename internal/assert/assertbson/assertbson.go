// Copyright (C) MongoDB, Inc. 2025-present.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at http://www.apache.org/licenses/LICENSE-2.0

package assertbson

import (
	"github.com/hanzoai/documentdb-go/bson"
	"github.com/hanzoai/documentdb-go/internal/assert"
	"github.com/hanzoai/documentdb-go/x/bsonx/bsoncore"
)

type tHelper interface {
	Helper()
}

// EqualDocument asserts that the expected and actual BSON documents are equal.
// If the documents are not equal, it prints both the binary diff and Extended
// JSON representation of the BSON documents.
func EqualDocument(t assert.TestingT, expected, actual []byte) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	return assert.Equal(t,
		expected,
		actual,
		`expected and actual BSON documents do not match
As Extended JSON:
Expected: %s
Actual  : %s`,
		bson.Raw(expected),
		bson.Raw(actual))
}

// EqualValue asserts that the expected and actual BSON values are equal. If the
// values are not equal, it prints both the binary diff and Extended JSON
// representation of the BSON values.
func EqualValue[T bson.RawValue | bsoncore.Value](t assert.TestingT, expected, actual T) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	return assert.Equal(t,
		expected,
		actual,
		`expected and actual BSON values do not match
As Extended JSON:
Expected: %s
Actual  : %s`,
		expected,
		actual)
}
