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
	"sync"
)

// Component represents the service/component that owns the error code
type Component int

const (
	ComponentInfra Component = iota
	ComponentFrontend
	ComponentHistory
	ComponentMatching
	ComponentWorker
	ComponentWorkflow
	ComponentSystem
	ComponentExternal
	ComponentResource
	ComponentCommon
	ComponentPersist
	ComponentTools
	ComponentSchema
	ComponentArchiver
	ComponentTest
)

// String returns the string representation of the Component
func (c Component) String() string {
	switch c {
	case ComponentInfra:
		return "INFRA"
	case ComponentFrontend:
		return "FRONT"
	case ComponentHistory:
		return "HIST"
	case ComponentMatching:
		return "MATCH"
	case ComponentWorker:
		return "WORK"
	case ComponentWorkflow:
		return "WF"
	case ComponentSystem:
		return "SYS"
	case ComponentExternal:
		return "EXT"
	case ComponentResource:
		return "RES"
	case ComponentCommon:
		return "COMMON"
	case ComponentPersist:
		return "PERSIST"
	case ComponentTools:
		return "TOOLS"
	case ComponentSchema:
		return "SCHEMA"
	case ComponentArchiver:
		return "ARCHIVER"
	case ComponentTest:
		return "TEST"
	default:
		return "UNKNOWN"
	}
}

// GetRange returns the valid error code range for the component
func (c Component) GetRange() (int, int) {
	switch c {
	case ComponentInfra:
		return 1000, 1999
	case ComponentFrontend:
		return 2000, 2999
	case ComponentHistory:
		return 3000, 3999
	case ComponentMatching:
		return 4000, 4999
	case ComponentWorker:
		return 5000, 5999
	case ComponentWorkflow:
		return 6000, 6999
	case ComponentSystem:
		return 7000, 7999
	case ComponentExternal:
		return 8000, 8999
	case ComponentResource:
		return 9000, 9999
	case ComponentCommon:
		return 7000, 9099 // Common spans multiple ranges
	case ComponentPersist:
		return 7500, 9099 // Persistence spans multiple ranges
	case ComponentTools:
		return 7200, 7299
	case ComponentSchema:
		return 8200, 8299
	case ComponentArchiver:
		return 7100, 9099 // Archiver spans multiple ranges
	case ComponentTest:
		return 7600, 7699
	default:
		return 0, 0
	}
}

type ErrorCode struct {
	Code      int
	Component Component
	Message   string
	Severity  string
	Retryable bool
}

type Registry struct {
	codes map[int]ErrorCode
	mutex sync.RWMutex
}

var globalRegistry = NewRegistry()

func NewRegistry() *Registry {
	return &Registry{
		codes: make(map[int]ErrorCode),
	}
}

// Register registers an error code with the given component enum
func Register(code int, component Component, message string, opts ...Option) error {
	globalRegistry.mutex.Lock()
	defer globalRegistry.mutex.Unlock()

	if _, exists := globalRegistry.codes[code]; exists {
		return fmt.Errorf("error code %d already registered", code)
	}

	ec := ErrorCode{
		Code:      code,
		Component: component,
		Message:   message,
		Severity:  "ERROR", // default
		Retryable: false,   // default
	}

	for _, opt := range opts {
		opt(&ec)
	}

	// Validate code is in correct range for component
	if err := validateCodeRange(code, component); err != nil {
		return err
	}

	globalRegistry.codes[code] = ec
	return nil
}

// RegisterString is a backward-compatible wrapper that accepts string components
func RegisterString(code int, componentStr, message string, opts ...Option) error {
	component := parseComponentString(componentStr)
	return Register(code, component, message, opts...)
}

// parseComponentString converts string component names to Component enum
func parseComponentString(componentStr string) Component {
	switch componentStr {
	case "INFRA":
		return ComponentInfra
	case "FRONT":
		return ComponentFrontend
	case "HIST":
		return ComponentHistory
	case "MATCH":
		return ComponentMatching
	case "WORK":
		return ComponentWorker
	case "WF":
		return ComponentWorkflow
	case "SYS":
		return ComponentSystem
	case "EXT":
		return ComponentExternal
	case "RES":
		return ComponentResource
	case "COMMON":
		return ComponentCommon
	case "PERSIST":
		return ComponentPersist
	case "TOOLS":
		return ComponentTools
	case "SCHEMA":
		return ComponentSchema
	case "ARCHIVER":
		return ComponentArchiver
	case "TEST":
		return ComponentTest
	default:
		return ComponentCommon // Default fallback
	}
}

func Get(code int) (ErrorCode, bool) {
	globalRegistry.mutex.RLock()
	defer globalRegistry.mutex.RUnlock()

	ec, exists := globalRegistry.codes[code]
	return ec, exists
}

func GetAllCodes() []ErrorCode {
	globalRegistry.mutex.RLock()
	defer globalRegistry.mutex.RUnlock()

	codes := make([]ErrorCode, 0, len(globalRegistry.codes))
	for _, code := range globalRegistry.codes {
		codes = append(codes, code)
	}
	return codes
}

type Option func(*ErrorCode)

func WithSeverity(severity string) Option {
	return func(ec *ErrorCode) {
		ec.Severity = severity
	}
}

func WithRetryable(retryable bool) Option {
	return func(ec *ErrorCode) {
		ec.Retryable = retryable
	}
}

func validateCodeRange(code int, component Component) error {
	minRange, maxRange := component.GetRange()
	
	if minRange == 0 && maxRange == 0 {
		return fmt.Errorf("invalid component: %v", component)
	}

	if code < minRange || code > maxRange {
		return fmt.Errorf("code %d not in range %d-%d for component %s",
			code, minRange, maxRange, component.String())
	}

	return nil
}
