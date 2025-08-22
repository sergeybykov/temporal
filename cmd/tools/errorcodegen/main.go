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

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		output = flag.String("output", "", "Output file path for generated constants")
		verify = flag.Bool("verify", false, "Verify no conflicts in existing codes")
		scan   = flag.Bool("scan", false, "Scan codebase and generate error codes")
	)
	flag.Parse()

	if *verify {
		if err := verifyErrorCodes(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Error code verification passed")
		return
	}

	if *scan {
		if *output == "" {
			log.Fatal("Output file required for scan operation")
		}
		if err := scanAndGenerateErrorCodes(*output); err != nil {
			log.Fatal(err)
		}
		return
	}

	if *output == "" {
		log.Fatal("Output file required")
	}

	if err := generateErrorCodes(*output); err != nil {
		log.Fatal(err)
	}
}

// New automated scanning approach
func scanAndGenerateErrorCodes(outputPath string) error {
	fmt.Println("Scanning codebase for error logging patterns...")

	// Scan codebase for error logging patterns
	errorLocations, err := scanCodebaseForErrors()
	if err != nil {
		return fmt.Errorf("failed to scan codebase: %w", err)
	}

	fmt.Printf("Found %d error logging locations\n", len(errorLocations))

	// Generate sequential error codes for each location
	codeAssignments, err := assignSequentialErrorCodes(errorLocations)
	if err != nil {
		return fmt.Errorf("failed to assign error codes: %w", err)
	}

	// Generate constants file
	return generateConstantsFile(outputPath, codeAssignments)
}

type ErrorLocation struct {
	FilePath   string
	LineNumber int
	Function   string
	Message    string
	Component  string
	LogLevel   string
}

func scanCodebaseForErrors() ([]ErrorLocation, error) {
	var locations []ErrorLocation

	// Define source directories to scan
	sourceDirs := []string{
		"service/frontend",
		"service/history",
		"service/matching",
		"service/worker",
		"common/persistence",
		"common/log",
		"client",
	}

	for _, dir := range sourceDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			fmt.Printf("Skipping non-existent directory: %s\n", dir)
			continue
		}

		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Only scan Go source files, skip tests
			if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
				return nil
			}

			fileLocations, err := scanFileForErrors(path)
			if err != nil {
				fmt.Printf("Warning: failed to scan file %s: %v\n", path, err)
				return nil // Continue scanning other files
			}

			locations = append(locations, fileLocations...)
			return nil
		})

		if err != nil {
			return nil, fmt.Errorf("failed to walk directory %s: %w", dir, err)
		}
	}

	return locations, nil
}

func scanFileForErrors(filePath string) ([]ErrorLocation, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var locations []ErrorLocation
	lines := strings.Split(string(content), "\n")

	// Enhanced patterns to match various error logging calls
	errorPatterns := []*regexp.Regexp{
		regexp.MustCompile(`\.Error\s*\(`),
		regexp.MustCompile(`\.Errorf\s*\(`),
		regexp.MustCompile(`\.Warn\s*\(`),
		regexp.MustCompile(`\.Warnf\s*\(`),
		regexp.MustCompile(`\.Fatal\s*\(`),
		regexp.MustCompile(`\.Fatalf\s*\(`),
		regexp.MustCompile(`logger\.Error`),
		regexp.MustCompile(`logger\.Warn`),
		regexp.MustCompile(`logger\.Fatal`),
	}

	for lineNum, line := range lines {
		// Skip comments and empty lines
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") || trimmed == "" {
			continue
		}

		for _, pattern := range errorPatterns {
			if pattern.MatchString(line) {
				location := ErrorLocation{
					FilePath:   filePath,
					LineNumber: lineNum + 1,
					Message:    extractErrorMessage(line),
					Component:  determineComponent(filePath),
					LogLevel:   extractLogLevel(line),
				}

				locations = append(locations, location)
				break // Only match once per line
			}
		}
	}

	return locations, nil
}

func determineComponent(filePath string) string {
	switch {
	case strings.Contains(filePath, "service/frontend"):
		return "FRONT"
	case strings.Contains(filePath, "service/history"):
		return "HIST"
	case strings.Contains(filePath, "service/matching"):
		return "MATCH"
	case strings.Contains(filePath, "service/worker"):
		return "WORK"
	case strings.Contains(filePath, "common/persistence"):
		return "INFRA"
	case strings.Contains(filePath, "common"):
		return "SYS"
	case strings.Contains(filePath, "client"):
		return "SYS"
	default:
		return "SYS" // default
	}
}

func extractErrorMessage(line string) string {
	// Extract first string literal from logging call
	re := regexp.MustCompile(`"([^"]*)"`)
	matches := re.FindStringSubmatch(line)
	if len(matches) > 1 && matches[1] != "" {
		return matches[1]
	}
	return "error logging call"
}

func extractLogLevel(line string) string {
	switch {
	case strings.Contains(line, ".Error") || strings.Contains(line, "logger.Error"):
		return "ERROR"
	case strings.Contains(line, ".Warn") || strings.Contains(line, "logger.Warn"):
		return "WARN"
	case strings.Contains(line, ".Fatal") || strings.Contains(line, "logger.Fatal"):
		return "FATAL"
	default:
		return "ERROR"
	}
}

type CodeAssignment struct {
	Location ErrorLocation
	Code     int
	Constant string
}

func assignSequentialErrorCodes(locations []ErrorLocation) ([]CodeAssignment, error) {
	// Group locations by component and assign sequential codes
	componentRanges := map[string][2]int{
		"INFRA": {1000, 1999},
		"FRONT": {2000, 2999},
		"HIST":  {3000, 3999},
		"MATCH": {4000, 4999},
		"WORK":  {5000, 5999},
		"WF":    {6000, 6999},
		"SYS":   {7000, 7999},
		"EXT":   {8000, 8999},
	}

	// Sort locations for consistent assignment - simple file path and line number ordering
	sort.Slice(locations, func(i, j int) bool {
		if locations[i].Component != locations[j].Component {
			return locations[i].Component < locations[j].Component
		}
		if locations[i].FilePath != locations[j].FilePath {
			return locations[i].FilePath < locations[j].FilePath
		}
		return locations[i].LineNumber < locations[j].LineNumber
	})

	componentCounters := make(map[string]int)
	var assignments []CodeAssignment

	for _, location := range locations {
		component := location.Component
		rangeStart := componentRanges[component][0]

		if _, exists := componentCounters[component]; !exists {
			componentCounters[component] = rangeStart
		}

		code := componentCounters[component]
		componentCounters[component]++

		// Check if we exceed the range
		if code > componentRanges[component][1] {
			return nil, fmt.Errorf("exceeded error code range for component %s at code %d", component, code)
		}

		// Generate simple sequential constant name
		constantName := fmt.Sprintf("%s_%04d", component, code)

		assignments = append(assignments, CodeAssignment{
			Location: location,
			Code:     code,
			Constant: constantName,
		})
	}

	return assignments, nil
}

func generateConstantsFile(outputPath string, assignments []CodeAssignment) error {
	var buf strings.Builder

	buf.WriteString(`// Code generated by errorcodegen. DO NOT EDIT.
package errorcode

// Error code constants - each represents a unique error logging location
const (
`)

	// Group by component for better organization
	componentGroups := make(map[string][]CodeAssignment)
	for _, assignment := range assignments {
		comp := assignment.Location.Component
		componentGroups[comp] = append(componentGroups[comp], assignment)
	}

	// Generate constants grouped by component
	componentOrder := []string{"INFRA", "FRONT", "HIST", "MATCH", "WORK", "WF", "SYS", "EXT"}

	for _, component := range componentOrder {
		assignments, exists := componentGroups[component]
		if !exists || len(assignments) == 0 {
			continue
		}

		buf.WriteString(fmt.Sprintf("\n\t// %s Service Error Codes (%d codes)\n", component, len(assignments)))

		for _, assignment := range assignments {
			comment := fmt.Sprintf("%s:%d",
				strings.TrimPrefix(assignment.Location.FilePath, "./"),
				assignment.Location.LineNumber)
			if assignment.Location.Message != "" && assignment.Location.Message != "error logging call" {
				comment += fmt.Sprintf(" - %s", assignment.Location.Message)
			}

			buf.WriteString(fmt.Sprintf("\t%-40s = %d // %s\n",
				assignment.Constant,
				assignment.Code,
				comment,
			))
		}
	}

	buf.WriteString(")\n\n")

	// Add registration function
	buf.WriteString(`func init() {
	// Auto-register all generated error codes
`)

	for _, component := range componentOrder {
		assignments, exists := componentGroups[component]
		if !exists || len(assignments) == 0 {
			continue
		}

		buf.WriteString(fmt.Sprintf("\tregister%sCodes()\n", component))
	}

	buf.WriteString("}\n\n")

	// Generate registration functions
	for _, component := range componentOrder {
		assignments, exists := componentGroups[component]
		if !exists || len(assignments) == 0 {
			continue
		}

		buf.WriteString(fmt.Sprintf("func register%sCodes() {\n", component))
		for _, assignment := range assignments {
			buf.WriteString(fmt.Sprintf("\tRegister(%s, \"%s\", \"%s\")\n",
				assignment.Constant,
				assignment.Location.Component,
				assignment.Location.Message,
			))
		}
		buf.WriteString("}\n\n")
	}

	fmt.Printf("Generated %d error codes in %s\n", len(assignments), outputPath)
	return os.WriteFile(outputPath, []byte(buf.String()), 0644)
}

// Legacy manual generation function
func generateErrorCodes(outputPath string) error {
	// Implement manual generation as fallback
	return fmt.Errorf("manual generation not implemented - use -scan flag for automated generation")
}

func verifyErrorCodes() error {
	// Parse error code constants and verify ranges
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "common/errorcode/codes.go", nil, parser.ParseComments)
	if err != nil {
		return err
	}

	codes := make(map[int]string)
	ranges := map[string][2]int{
		"FRONT": {2000, 2999},
		"HIST":  {3000, 3999},
		"MATCH": {4000, 4999},
		"WORK":  {5000, 5999},
		"INFRA": {1000, 1999},
		"SYS":   {7000, 7999},
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.CONST {
			for _, spec := range decl.Specs {
				if vspec, ok := spec.(*ast.ValueSpec); ok {
					for i, name := range vspec.Names {
						if i < len(vspec.Values) {
							if lit, ok := vspec.Values[i].(*ast.BasicLit); ok && lit.Kind == token.INT {
								code, _ := strconv.Atoi(lit.Value)
								if existing, exists := codes[code]; exists {
									fmt.Printf("ERROR: duplicate error code %d: %s and %s\n", code, existing, name.Name)
									return false
								}
								codes[code] = name.Name

								// Verify range
								component := getComponentFromName(name.Name)
								if r, exists := ranges[component]; exists {
									if code < r[0] || code > r[1] {
										fmt.Printf("ERROR: error code %d (%s) not in range %d-%d for component %s\n",
											code, name.Name, r[0], r[1], component)
										return false
									}
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	return nil
}

func getComponentFromName(name string) string {
	switch {
	case strings.HasPrefix(name, "FRONT"):
		return "FRONT"
	case strings.HasPrefix(name, "HIST"):
		return "HIST"
	case strings.HasPrefix(name, "MATCH"):
		return "MATCH"
	case strings.HasPrefix(name, "WORK"):
		return "WORK"
	case strings.HasPrefix(name, "INFRA"):
		return "INFRA"
	case strings.HasPrefix(name, "SYS"):
		return "SYS"
	default:
		return "UNKNOWN"
	}
}
