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

package log

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/server/common/errorcode"
	"go.temporal.io/server/common/log/tag"
)

func TestErrorCodeTagFormatting(t *testing.T) {
	// Test the error code tag formatting
	errorCodeTag := tag.ErrorCode(2001)

	require.Equal(t, "error-code", errorCodeTag.Key())
	require.Equal(t, int64(2001), errorCodeTag.Value()) // NewInt returns int64 when accessed via Value()
}

func TestErrorComponentTag(t *testing.T) {
	// Test the error component tag
	componentTag := tag.ErrorComponent("FRONT")

	require.Equal(t, "error-component", componentTag.Key())
	require.Equal(t, "FRONT", componentTag.Value())
}

func TestErrorSeverityTag(t *testing.T) {
	// Test the error severity tag
	severityTag := tag.ErrorSeverity("ERROR")

	require.Equal(t, "error-severity", severityTag.Key())
	require.Equal(t, "ERROR", severityTag.Value())
}

func TestErrorWithCodeBasic(t *testing.T) {
	// Test that ErrorWithCode function can be called without errors
	// We're just testing that the function signatures are correct and no panics occur

	// Use a simple test logger that captures calls
	testLogger := &testCaptureLogger{}

	ErrorWithCode(testLogger, errorcode.FrontendInvalidWorkflowID, "test message", nil)

	require.True(t, testLogger.errorCalled)
	require.Equal(t, "test message", testLogger.lastMessage)
	require.True(t, len(testLogger.lastTags) >= 3) // Should have at least error-code, component, severity
}

// Simple test logger that captures the last call
type testCaptureLogger struct {
	errorCalled bool
	warnCalled  bool
	fatalCalled bool
	lastMessage string
	lastTags    []tag.Tag
}

func (t *testCaptureLogger) Debug(msg string, tags ...tag.Tag) {
	t.lastMessage = msg
	t.lastTags = tags
}

func (t *testCaptureLogger) Info(msg string, tags ...tag.Tag) {
	t.lastMessage = msg
	t.lastTags = tags
}

func (t *testCaptureLogger) Warn(msg string, tags ...tag.Tag) {
	t.warnCalled = true
	t.lastMessage = msg
	t.lastTags = tags
}

func (t *testCaptureLogger) Error(msg string, tags ...tag.Tag) {
	t.errorCalled = true
	t.lastMessage = msg
	t.lastTags = tags
}

func (t *testCaptureLogger) DPanic(msg string, tags ...tag.Tag) {
	t.lastMessage = msg
	t.lastTags = tags
}

func (t *testCaptureLogger) Panic(msg string, tags ...tag.Tag) {
	t.lastMessage = msg
	t.lastTags = tags
}

func (t *testCaptureLogger) Fatal(msg string, tags ...tag.Tag) {
	t.fatalCalled = true
	t.lastMessage = msg
	t.lastTags = tags
}

func TestWarnWithCode(t *testing.T) {
	testLogger := &testCaptureLogger{}

	WarnWithCode(testLogger, errorcode.FrontendInvalidWorkflowID, "test warning message")

	require.True(t, testLogger.warnCalled)
	require.Equal(t, "test warning message", testLogger.lastMessage)
	require.True(t, len(testLogger.lastTags) >= 3) // Should have error-code, component, severity
}

func TestFatalWithCode(t *testing.T) {
	testLogger := &testCaptureLogger{}
	testErr := errors.New("test error")

	FatalWithCode(testLogger, errorcode.FrontendInvalidWorkflowID, "fatal error occurred", testErr)

	require.True(t, testLogger.fatalCalled)
	require.Equal(t, "fatal error occurred", testLogger.lastMessage)
	require.True(t, len(testLogger.lastTags) >= 4) // Should have error-code, component, severity, error
}

func TestErrorWithCodeInterface(t *testing.T) {
	// Test the ErrorWithCodeInterface implementation
	testErr := &codedError{
		code:    1001,
		message: "test coded error",
		cause:   errors.New("underlying error"),
	}

	require.Equal(t, 1001, testErr.ErrorCode())
	require.Contains(t, testErr.Error(), "test coded error")
	require.Contains(t, testErr.Error(), "underlying error")

	// Test unwrapping
	unwrapped := errors.Unwrap(testErr)
	require.NotNil(t, unwrapped)
	require.Equal(t, "underlying error", unwrapped.Error())
}

func TestErrorWithCodeFromError(t *testing.T) {
	testLogger := &testCaptureLogger{}

	// Test with coded error
	codedErr := CreateCodedError(1001, "coded error message", nil)
	ErrorWithCodeFromError(testLogger, "processing failed", codedErr)

	require.True(t, testLogger.errorCalled)
	require.Equal(t, "processing failed", testLogger.lastMessage)

	// Reset logger
	testLogger.errorCalled = false
	testLogger.lastMessage = ""
	testLogger.lastTags = nil

	// Test with regular error
	regularErr := errors.New("regular error")
	ErrorWithCodeFromError(testLogger, "regular processing failed", regularErr)

	require.True(t, testLogger.errorCalled)
	require.Equal(t, "regular processing failed", testLogger.lastMessage)
}

func TestContextualErrorWithCode(t *testing.T) {
	testLogger := &testCaptureLogger{}

	// Create context with some values
	ctx := context.WithValue(context.Background(), "namespace", "test-namespace")
	ctx = context.WithValue(ctx, "workflowID", "test-workflow-123")

	ContextualErrorWithCode(ctx, testLogger, errorcode.FrontendInvalidWorkflowID, "contextual error", nil)

	require.True(t, testLogger.errorCalled)
	require.Equal(t, "contextual error", testLogger.lastMessage)
	require.True(t, len(testLogger.lastTags) >= 3) // Should have error-code, component, severity + context tags

	// Check if context tags are included
	hasNamespaceTag := false
	hasWorkflowIDTag := false
	for _, tag := range testLogger.lastTags {
		if tag.Key() == "namespace" && tag.Value() == "test-namespace" {
			hasNamespaceTag = true
		}
		if tag.Key() == "workflow-id" && tag.Value() == "test-workflow-123" {
			hasWorkflowIDTag = true
		}
	}
	require.True(t, hasNamespaceTag, "Should include namespace tag from context")
	require.True(t, hasWorkflowIDTag, "Should include workflow ID tag from context")
}

func TestContextualWarnWithCode(t *testing.T) {
	testLogger := &testCaptureLogger{}

	ctx := context.WithValue(context.Background(), "namespace", "test-namespace")

	ContextualWarnWithCode(ctx, testLogger, errorcode.FrontendInvalidWorkflowID, "contextual warning")

	require.True(t, testLogger.warnCalled)
	require.Equal(t, "contextual warning", testLogger.lastMessage)
	require.True(t, len(testLogger.lastTags) >= 3)
}

func TestCreateCodedError(t *testing.T) {
	// Test creating coded error without cause
	err1 := CreateCodedError(1001, "test error", nil)
	require.Equal(t, "test error", err1.Error())

	// Test with cause
	cause := errors.New("underlying cause")
	err2 := CreateCodedError(1002, "wrapper error", cause)
	require.Contains(t, err2.Error(), "wrapper error")
	require.Contains(t, err2.Error(), "underlying cause")

	// Test ErrorWithCodeInterface
	codedErr, ok := err2.(ErrorWithCodeInterface)
	require.True(t, ok)
	require.Equal(t, 1002, codedErr.ErrorCode())
}

func TestWrapErrorWithCode(t *testing.T) {
	// Test wrapping nil
	err1 := WrapErrorWithCode(1001, "wrapper message", nil)
	require.Equal(t, "wrapper message", err1.Error())

	// Test wrapping existing error
	cause := errors.New("original error")
	err2 := WrapErrorWithCode(1002, "wrapper message", cause)
	require.Contains(t, err2.Error(), "wrapper message")
	require.Contains(t, err2.Error(), "original error")

	// Test ErrorCode extraction
	code, ok := GetErrorCodeFromError(err2)
	require.True(t, ok)
	require.Equal(t, 1002, code)
}

func TestGetErrorCodeFromError(t *testing.T) {
	// Test with coded error
	codedErr := CreateCodedError(1001, "coded error", nil)
	code, ok := GetErrorCodeFromError(codedErr)
	require.True(t, ok)
	require.Equal(t, 1001, code)

	// Test with regular error
	regularErr := errors.New("regular error")
	code, ok = GetErrorCodeFromError(regularErr)
	require.False(t, ok)
	require.Equal(t, 0, code)
}
