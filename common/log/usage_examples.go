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

// Package log provides examples of how to use the enhanced error code logging system.
// This file contains usage examples and is not meant to be imported or used in production.
//
//go:build examples
// +build examples

package log

import (
	"context"
	"errors"

	"go.temporal.io/server/common/errorcode"
	"go.temporal.io/server/common/log/tag"
)

// ExampleBasicErrorLogging demonstrates basic error code logging
func ExampleBasicErrorLogging() {
	var logger Logger
	var err error

	// Basic error logging with code
	ErrorWithCode(logger, errorcode.FRONT_2001, "Invalid workflow ID provided", err,
		tag.WorkflowID("invalid-wf-id"),
		tag.Operation("StartWorkflow"))

	// Warning with code
	WarnWithCode(logger, errorcode.HIST_3001, "Workflow execution taking longer than expected",
		tag.WorkflowID("slow-workflow"),
		tag.Duration("5m"))

	// Fatal error with code
	FatalWithCode(logger, errorcode.INFRA_1001, "Cannot connect to database", err,
		tag.DatabaseName("temporal"),
		tag.Address("localhost:5432"))
}

// ExampleContextualLogging demonstrates contextual logging with error codes
func ExampleContextualLogging() {
	var logger Logger
	var err error

	// Create context with workflow information
	ctx := context.WithValue(context.Background(), "namespace", "my-namespace")
	ctx = context.WithValue(ctx, "workflowID", "my-workflow-123")
	ctx = context.WithValue(ctx, "runID", "run-456")

	// Contextual error logging - automatically extracts context information
	ContextualErrorWithCode(ctx, logger, errorcode.HIST_3002, "Activity failed to complete", err,
		tag.ActivityName("ProcessOrder"),
		tag.ActivityID("activity-789"))

	// Contextual warning
	ContextualWarnWithCode(ctx, logger, errorcode.HIST_3003, "Activity retry attempt",
		tag.Attempt(3),
		tag.ActivityName("ProcessOrder"))
}

// ExampleErrorWithCodeInterface demonstrates using the ErrorWithCodeInterface
func ExampleErrorWithCodeInterface() {
	var logger Logger

	// Create a coded error
	codedErr := CreateCodedError(errorcode.FRONT_2002, "Workflow not found",
		errors.New("workflow with ID 'missing-wf' does not exist"))

	// Log using the ErrorWithCodeFromError helper
	ErrorWithCodeFromError(logger, "Failed to retrieve workflow", codedErr,
		tag.WorkflowID("missing-wf"))

	// Wrap an existing error with a code
	originalErr := errors.New("connection timeout")
	wrappedErr := WrapErrorWithCode(errorcode.INFRA_1002, "Database operation failed", originalErr)

	// Log the wrapped error
	ErrorWithCodeFromError(logger, "Could not save workflow state", wrappedErr,
		tag.Operation("SaveState"))
}

// ExampleMigrationScenario demonstrates migration from existing logging to error codes
func ExampleMigrationScenario() {
	var logger Logger
	var err error
	filename := "service/frontend/handler.go"
	line := 123

	// Migration-friendly logging - uses error codes if available, falls back otherwise
	LogErrorWithFallback(logger, filename, line, "Request validation failed", err,
		tag.RequestID("req-456"),
		tag.Operation("ValidateRequest"))

	// Migration with context
	ctx := context.WithValue(context.Background(), "namespace", "production")
	ContextualMigrateLoggerCall(ctx, logger, "ERROR", filename, line+10, "Processing error", err,
		tag.WorkflowID("prod-workflow"))
}

// ExampleServiceSpecificLogging shows logging patterns specific to different services
func ExampleServiceSpecificLogging() {
	var logger Logger
	var err error

	// Frontend service errors
	ErrorWithCode(logger, errorcode.FRONT_2100, "Authentication failed", err,
		tag.UserID("user-123"),
		tag.ClientName("temporal-cli"))

	// History service errors
	ErrorWithCode(logger, errorcode.HIST_3100, "Event cannot be applied to current state", err,
		tag.WorkflowID("workflow-456"),
		tag.EventID(42),
		tag.EventType("WorkflowTaskCompleted"))

	// Matching service errors
	ErrorWithCode(logger, errorcode.MATCH_4100, "No available workers for task queue", err,
		tag.TaskQueue("critical-tasks"),
		tag.TaskType("Activity"))

	// Worker service errors
	ErrorWithCode(logger, errorcode.WORK_5100, "Archival operation failed", err,
		tag.ArchivalURI("s3://bucket/path"),
		tag.Operation("Archive"))

	// Infrastructure errors
	ErrorWithCode(logger, errorcode.INFRA_1100, "Shard ownership lost", err,
		tag.ShardID(123),
		tag.Address("192.168.1.100"))
}

// ExampleErrorCodeExtractionAndHandling demonstrates extracting and handling error codes
func ExampleErrorCodeExtractionAndHandling() {
	var logger Logger

	// Create various coded errors
	err1 := CreateCodedError(errorcode.FRONT_2001, "Invalid input", nil)
	err2 := WrapErrorWithCode(errorcode.HIST_3001, "State transition error",
		errors.New("invalid state"))

	// Extract error codes
	if code, ok := GetErrorCodeFromError(err1); ok {
		logger.Info("Handling error with code")
	}

	// Handle multiple error types
	handleError := func(err error) {
		if code, ok := GetErrorCodeFromError(err); ok {
			// Handle based on error code
			switch {
			case code >= 2000 && code < 3000:
				logger.Info("Frontend error detected")
			case code >= 3000 && code < 4000:
				logger.Info("History service error detected")
			default:
				logger.Info("Other service error")
			}
		} else {
			// Handle legacy error without code
			ErrorWithCode(logger, errorcode.CommonLogMigrationOperationFailed, "Legacy error without code", err)
		}
	}

	handleError(err1)
	handleError(err2)
}
