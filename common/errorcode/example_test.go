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
	"fmt"
)

// Example demonstrates how to use error codes
func Example() {
	// Example: Use error code constants directly
	if ec, exists := Get(HistoryWorkflowNotFound); exists {
		fmt.Printf("Error Code: %d, Component: %s, Message: %s\n", ec.Code, ec.Component, ec.Message)
	}

	// Example: Check all available error codes
	allCodes := GetAllCodes()
	fmt.Printf("Total registered error codes: %d\n", len(allCodes))

	// Example: Show component ranges in order
	components := []struct{ name, desc string }{
		{"INFRA", "Infrastructure components (1000-1999)"},
		{"FRONT", "Frontend service (2000-2999)"},
		{"HIST", "History service (3000-3999)"},
		{"MATCH", "Matching service (4000-4999)"},
		{"WORK", "Worker service (5000-5999)"},
	}

	for _, comp := range components {
		fmt.Printf("Component %s: %s\n", comp.name, comp.desc)
	}

	// Output:
	// Error Code: 3001, Component: HIST, Message: Workflow execution not found
	// Total registered error codes: 227
	// Component INFRA: Infrastructure components (1000-1999)
	// Component FRONT: Frontend service (2000-2999)
	// Component HIST: History service (3000-3999)
	// Component MATCH: Matching service (4000-4999)
	// Component WORK: Worker service (5000-5999)
}
