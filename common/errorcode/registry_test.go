// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package errorcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrorCodeRegistration(t *testing.T) {
	codes := GetAllCodes()

	// Test for duplicates
	seen := make(map[int]bool)
	for _, code := range codes {
		require.False(t, seen[code.Code], "Duplicate error code: %d", code.Code)
		seen[code.Code] = true
	}

	// Test ranges
	for _, code := range codes {
		require.True(t, isValidRangeForComponent(code.Code, code.Component),
			"Code %d not in valid range for component %s", code.Code, code.Component.String())
	}
}

func TestErrorCodeConstants(t *testing.T) {
	// Test that constants match registered codes
	testCases := []struct {
		constant int
		expected string
	}{
		{FrontendInvalidWorkflowID, "FRONT"},
		{HistoryWorkflowNotFound, "HIST"},
		{MatchingTaskQueueNotFound, "MATCH"},
		{WorkerStartupFailed, "WORK"},
		{InfraDBConnectionFailed, "INFRA"},
	}

	for _, tc := range testCases {
		ec, exists := Get(tc.constant)
		require.True(t, exists, "Error code %d not registered", tc.constant)
		require.Equal(t, tc.expected, ec.Component.String())
	}
}

func TestRegisterDuplicate(t *testing.T) {
	// Create a temporary registry for testing
	testRegistry := NewRegistry()
	oldRegistry := globalRegistry
	globalRegistry = testRegistry
	defer func() { globalRegistry = oldRegistry }()

	// Register a code first
	err := Register(9999, ComponentResource, "Test message")
	require.NoError(t, err)

	// Try to register the same code again
	err = Register(9999, ComponentResource, "Duplicate test")
	require.Error(t, err)
	require.Contains(t, err.Error(), "already registered")
}

func TestInvalidComponent(t *testing.T) {
	// Test with an invalid component value (using a value outside the enum range)
	err := Register(1000, Component(999), "Test message")
	require.Error(t, err)
	require.Contains(t, err.Error(), "invalid component")
}

func TestInvalidRange(t *testing.T) {
	err := Register(500, ComponentFrontend, "Test message")
	require.Error(t, err)
	require.Contains(t, err.Error(), "not in range")
}

func TestWithOptions(t *testing.T) {
	// Create a temporary registry for testing
	testRegistry := NewRegistry()
	oldRegistry := globalRegistry
	globalRegistry = testRegistry
	defer func() { globalRegistry = oldRegistry }()

	err := Register(9500, ComponentResource, "Test message", WithSeverity("WARN"), WithRetryable(true))
	require.NoError(t, err)

	ec, exists := Get(9500)
	require.True(t, exists)
	require.Equal(t, "WARN", ec.Severity)
	require.True(t, ec.Retryable)
}

func isValidRangeForComponent(code int, component Component) bool {
	minRange, maxRange := component.GetRange()
	if minRange == 0 && maxRange == 0 {
		return false
	}
	return code >= minRange && code <= maxRange
}
