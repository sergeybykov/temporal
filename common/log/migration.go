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
	"fmt"
	"regexp"
	"strings"

	"go.temporal.io/server/common/errorcode"
	"go.temporal.io/server/common/log/tag"
)

// MigrationHelper provides utilities to help migrate existing logging calls to error-code-based logging
type MigrationHelper struct {
	codeMap map[string]int // Map from file:line to error code
}

// NewMigrationHelper creates a new migration helper
func NewMigrationHelper() *MigrationHelper {
	helper := &MigrationHelper{
		codeMap: make(map[string]int),
	}
	helper.loadGeneratedCodes()
	return helper
}

// loadGeneratedCodes loads the generated error codes mapping
func (m *MigrationHelper) loadGeneratedCodes() {
	// For now, this is a placeholder. In practice, this would parse
	// the generated codes file to build the mapping from file:line to error code
	// This functionality would be implemented when the migration tool is needed
}

// GetErrorCode returns the error code for a given file and line number
func (m *MigrationHelper) GetErrorCode(filename string, line int) (int, bool) {
	key := fmt.Sprintf("%s:%d", filename, line)
	code, exists := m.codeMap[key]
	return code, exists
}

// MigrateLoggerCall creates a migration-friendly logging call
// This function can be used as a drop-in replacement during migration
func MigrateLoggerCall(logger Logger, level string, filename string, line int, message string, err error, tags ...tag.Tag) {
	helper := NewMigrationHelper()
	if code, exists := helper.GetErrorCode(filename, line); exists {
		// Use error code logging
		switch strings.ToUpper(level) {
		case "ERROR":
			ErrorWithCode(logger, code, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		case "WARN":
			WarnWithCode(logger, code, message, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		case "FATAL":
			FatalWithCode(logger, code, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		default:
			ErrorWithCode(logger, code, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		}
	} else {
		// Fallback to regular logging
		switch strings.ToUpper(level) {
		case "ERROR":
			if err != nil {
				allTags := append([]tag.Tag{tag.Error(err)}, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
				logger.Error(message, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			} else {
				logger.Error(message, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			}
		case "WARN":
			logger.Warn(message, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		case "FATAL":
			if err != nil {
				allTags := append([]tag.Tag{tag.Error(err)}, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
				logger.Fatal(message, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			} else {
				logger.Fatal(message, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			}
		default:
			if err != nil {
				allTags := append([]tag.Tag{tag.Error(err)}, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
				logger.Error(message, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			} else {
				logger.Error(message, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			}
		}
	}
}

// ContextualMigrateLoggerCall includes context support
func ContextualMigrateLoggerCall(ctx context.Context, logger Logger, level string, filename string, line int, message string, err error, tags ...tag.Tag) {
	helper := NewMigrationHelper()
	if code, exists := helper.GetErrorCode(filename, line); exists {
		// Use contextual error code logging
		switch strings.ToUpper(level) {
		case "ERROR":
			ContextualErrorWithCode(ctx, logger, code, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		case "WARN":
			ContextualWarnWithCode(ctx, logger, code, message, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		case "FATAL":
			// Fatal doesn't typically use context, fallback to regular
			FatalWithCode(logger, code, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		default:
			ContextualErrorWithCode(ctx, logger, code, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		}
	} else {
		// Fallback with context tags
		ctxTags := extractContextTags(ctx)
		allTags := append(tags, ctxTags...)

		switch strings.ToUpper(level) {
		case "ERROR":
			if err != nil {
				errorTags := append([]tag.Tag{tag.Error(err)}, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
				logger.Error(message, append(errorTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			} else {
				logger.Error(message, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			}
		case "WARN":
			logger.Warn(message, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
		case "FATAL":
			if err != nil {
				errorTags := append([]tag.Tag{tag.Error(err)}, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
				logger.Fatal(message, append(errorTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			} else {
				logger.Fatal(message, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			}
		default:
			if err != nil {
				errorTags := append([]tag.Tag{tag.Error(err)}, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
				logger.Error(message, append(errorTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			} else {
				logger.Error(message, append(allTags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
			}
		}
	}
}

// extractLocationFromComment extracts file:line location from generated code comment
func extractLocationFromComment(comment string) string {
	// Look for pattern like "filename.go:123"
	re := regexp.MustCompile(`([^/\s]+\.go):(\d+)`)
	matches := re.FindStringSubmatch(comment)
	if len(matches) >= 3 {
		return fmt.Sprintf("%s:%s", matches[1], matches[2])
	}
	return ""
}

// LogErrorWithFallback logs an error with code if available, otherwise falls back to regular logging
// This is the most commonly used migration function
func LogErrorWithFallback(logger Logger, filename string, line int, message string, err error, tags ...tag.Tag) {
	MigrateLoggerCall(logger, "ERROR", filename, line, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
}

// LogWarnWithFallback logs a warning with code if available, otherwise falls back to regular logging
func LogWarnWithFallback(logger Logger, filename string, line int, message string, tags ...tag.Tag) {
	MigrateLoggerCall(logger, "WARN", filename, line, message, nil, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
}

// LogFatalWithFallback logs a fatal error with code if available, otherwise falls back to regular logging
func LogFatalWithFallback(logger Logger, filename string, line int, message string, err error, tags ...tag.Tag) {
	MigrateLoggerCall(logger, "FATAL", filename, line, message, err, append(tags, tag.ErrorCode(errorcode.CommonLogMigrationOperationFailed))...)
}

// CreateCodedError creates an error that implements ErrorWithCode interface
func CreateCodedError(code int, message string, cause error) error {
	return &codedError{
		code:    code,
		message: message,
		cause:   cause,
	}
}

// codedError implements ErrorWithCodeInterface
type codedError struct {
	code    int
	message string
	cause   error
}

func (e *codedError) Error() string {
	if e.cause != nil {
		return fmt.Sprintf("%s: %v", e.message, e.cause)
	}
	return e.message
}

func (e *codedError) ErrorCode() int {
	return e.code
}

func (e *codedError) Unwrap() error {
	return e.cause
}

// WrapErrorWithCode wraps an existing error with an error code
func WrapErrorWithCode(code int, message string, err error) error {
	if err == nil {
		return CreateCodedError(code, message, nil)
	}
	return CreateCodedError(code, message, err)
}

// GetErrorCodeFromError extracts error code from error if it implements ErrorWithCodeInterface
func GetErrorCodeFromError(err error) (int, bool) {
	if codedErr, ok := err.(ErrorWithCodeInterface); ok {
		return codedErr.ErrorCode(), true
	}
	return 0, false
}
