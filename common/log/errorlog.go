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

	"go.temporal.io/server/common/errorcode"
	"go.temporal.io/server/common/log/tag"
)

// ErrorWithCode logs an error message with error code
func ErrorWithCode(logger Logger, code int, message string, err error, tags ...tag.Tag) {
	allTags := make([]tag.Tag, 0, len(tags)+4)
	allTags = append(allTags, tag.ErrorCode(code))

	if ec, exists := errorcode.Get(code); exists {
		allTags = append(allTags, tag.ErrorComponent(ec.Component.String()))
		allTags = append(allTags, tag.ErrorSeverity(ec.Severity))
	}

	if err != nil {
		allTags = append(allTags, tag.Error(err))
	}

	allTags = append(allTags, tags...)
	logger.Error(message, allTags...)
}

// WarnWithCode logs a warning message with error code
func WarnWithCode(logger Logger, code int, message string, tags ...tag.Tag) {
	allTags := make([]tag.Tag, 0, len(tags)+3)
	allTags = append(allTags, tag.ErrorCode(code))

	if ec, exists := errorcode.Get(code); exists {
		allTags = append(allTags, tag.ErrorComponent(ec.Component.String()))
		allTags = append(allTags, tag.ErrorSeverity(ec.Severity))
	}

	allTags = append(allTags, tags...)
	logger.Warn(message, allTags...)
}

// FatalWithCode logs a fatal message with error code
func FatalWithCode(logger Logger, code int, message string, err error, tags ...tag.Tag) {
	allTags := make([]tag.Tag, 0, len(tags)+4)
	allTags = append(allTags, tag.ErrorCode(code))

	if ec, exists := errorcode.Get(code); exists {
		allTags = append(allTags, tag.ErrorComponent(ec.Component.String()))
		allTags = append(allTags, tag.ErrorSeverity(ec.Severity))
	}

	if err != nil {
		allTags = append(allTags, tag.Error(err))
	}

	allTags = append(allTags, tags...)
	logger.Fatal(message, allTags...)
}

// ErrorWithCodeInterface interface for errors that contain error codes
type ErrorWithCodeInterface interface {
	error
	ErrorCode() int
}

// ErrorWithCodeFromError extracts error code from ErrorWithCodeInterface and logs appropriately
func ErrorWithCodeFromError(logger Logger, message string, err error, tags ...tag.Tag) {
	if codedErr, ok := err.(ErrorWithCodeInterface); ok {
		ErrorWithCode(logger, codedErr.ErrorCode(), message, err, tags...)
	} else {
		// Fallback to using a generic error code for uncoded errors
		ErrorWithCode(logger, errorcode.CommonLogMigrationOperationFailed, message, err, tags...)
	}
}

// ContextualErrorWithCode logs an error with code and context information
func ContextualErrorWithCode(ctx context.Context, logger Logger, code int, message string, err error, tags ...tag.Tag) {
	// Extract useful context information if available
	ctxTags := extractContextTags(ctx)
	allTags := append(tags, ctxTags...)
	ErrorWithCode(logger, code, message, err, allTags...)
}

// ContextualWarnWithCode logs a warning with code and context information
func ContextualWarnWithCode(ctx context.Context, logger Logger, code int, message string, tags ...tag.Tag) {
	// Extract useful context information if available
	ctxTags := extractContextTags(ctx)
	allTags := append(tags, ctxTags...)
	WarnWithCode(logger, code, message, allTags...)
}

// extractContextTags extracts useful logging tags from context
func extractContextTags(ctx context.Context) []tag.Tag {
	var tags []tag.Tag

	// Extract common context values that are useful for logging
	// This can be extended based on what context values are commonly used

	// Example: Extract namespace from context if available
	if ns := ctx.Value("namespace"); ns != nil {
		if nsStr, ok := ns.(string); ok {
			tags = append(tags, tag.NewStringTag("namespace", nsStr))
		}
	}

	// Example: Extract workflow ID from context if available
	if wfID := ctx.Value("workflowID"); wfID != nil {
		if wfIDStr, ok := wfID.(string); ok {
			tags = append(tags, tag.NewStringTag("workflow-id", wfIDStr))
		}
	}

	// Example: Extract run ID from context if available
	if runID := ctx.Value("runID"); runID != nil {
		if runIDStr, ok := runID.(string); ok {
			tags = append(tags, tag.NewStringTag("run-id", runIDStr))
		}
	}

	return tags
}
