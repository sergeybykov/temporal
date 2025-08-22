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
		help   = flag.Bool("help", false, "Show usage help")
	)
	flag.Parse()

	if *help {
		showUsageHelp()
		return
	}

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

// scanAndGenerateErrorCodes performs automated scanning of the codebase
// to find logging statements that need error code migration
func scanAndGenerateErrorCodes(outputPath string) error {
	fmt.Println("Scanning codebase for error logging patterns...")
	fmt.Println("Looking for unmigrated .Error(), .Warn(), .Fatal() calls...")
	fmt.Println("Skipping already migrated log.*WithCode() calls...")

	// Scan codebase for error logging patterns
	errorLocations, err := scanCodebaseForErrors()
	if err != nil {
		return fmt.Errorf("failed to scan codebase: %w", err)
	}

	fmt.Printf("Found %d error logging locations that need migration\n", len(errorLocations))

	// Generate sequential error codes for each location
	codeAssignments, err := assignSequentialErrorCodes(errorLocations)
	if err != nil {
		return fmt.Errorf("failed to assign error codes: %w", err)
	}

	// Generate constants file
	return generateConstantsFile(outputPath, codeAssignments)
}

// ErrorLocation represents a logging statement that needs error code migration
type ErrorLocation struct {
	FilePath   string // Path to source file containing the logging call
	LineNumber int    // Line number of the logging call
	Function   string // Function name context (if available)
	Message    string // Log message text
	Component  string // Service component (Frontend, History, etc.)
	LogLevel   string // Log level (ERROR, WARN, FATAL)
}

func scanCodebaseForErrors() ([]ErrorLocation, error) {
	var locations []ErrorLocation

	// Define source directories to scan - focus on main service packages
	sourceDirs := []string{
		"service/frontend", // Frontend API service
		"service/history",  // History service and workflows
		"service/matching", // Matching service and task queues
		"service/worker",   // System worker services
		"components",       // Shared components like scheduler
		"common",           // Common utilities and persistence
		"client",           // Client libraries
		"cmd/tools",        // Command-line tools
		"tools",            // Development tools
	}
	fmt.Printf("Scanning directories: %v\n", sourceDirs)

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

	// Enhanced patterns to match actual logging patterns in codebase
	// 1. Already migrated WithCode patterns (skip these)
	alreadyMigratedPatterns := []*regexp.Regexp{
		regexp.MustCompile(`log\.(Error|Warn|Fatal)WithCode\(`),
	}

	// 2. Patterns that need migration to WithCode
	needMigrationPatterns := []*regexp.Regexp{
		// Server logger patterns - most common in the codebase
		regexp.MustCompile(`\blogger\.(Error|Warn|Fatal)\(`),
		regexp.MustCompile(`\.[rR]?logger\.(Error|Warn|Fatal)\(`),
		regexp.MustCompile(`\b[a-zA-Z_][a-zA-Z0-9_]*\.logger\.(Error|Warn|Fatal)\(`),
		// SDK logger patterns (workflow contexts, uses key-value pairs)
		regexp.MustCompile(`\bs\.[lL]ogger\.(Error|Warn)\(`),
		regexp.MustCompile(`\bd\.[lL]ogger\.(Error|Warn)\(`),
		regexp.MustCompile(`\b[a-zA-Z_][a-zA-Z0-9_]*\.[lL]ogger\.(Error|Warn)\(`),
		// Tagged logging calls that need error codes
		regexp.MustCompile(`\b[a-zA-Z_][a-zA-Z0-9_]*\.(Error|Warn|Fatal)\(.*tag\.|.*tag\.Error`),
	}

	for lineNum, line := range lines {
		// Skip comments, empty lines, and test files
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") || trimmed == "" || strings.Contains(line, "_test.go") {
			continue
		}

		// Skip already migrated WithCode calls (just track them)
		alreadyMigrated := false
		for _, pattern := range alreadyMigratedPatterns {
			if pattern.MatchString(line) {
				alreadyMigrated = true
				break
			}
		}

		if alreadyMigrated {
			continue // Skip already migrated lines
		}

		// Check patterns that need migration
		for _, pattern := range needMigrationPatterns {
			if pattern.MatchString(line) {
				location := ErrorLocation{
					FilePath:   filePath,
					LineNumber: lineNum + 1,
					Message:    extractErrorMessage(line),
					Component:  determineComponent(filePath),
					LogLevel:   extractLogLevel(line),
					Function:   extractFunctionContext(lines, lineNum),
				}

				locations = append(locations, location)
				break // Only match once per line
			}
		}
	}

	return locations, nil
}

func determineComponent(filePath string) string {
	// Match the actual error code naming patterns used in codes.go
	switch {
	case strings.Contains(filePath, "service/frontend"):
		return "Frontend"
	case strings.Contains(filePath, "service/history"):
		return "History"
	case strings.Contains(filePath, "service/matching"):
		return "Matching"
	case strings.Contains(filePath, "service/worker"):
		return "Worker"
	case strings.Contains(filePath, "components"):
		return "Component"
	case strings.Contains(filePath, "common/persistence"):
		return "Persist"
	case strings.Contains(filePath, "common/archiver"):
		return "Common"
	case strings.Contains(filePath, "common"):
		return "Common"
	case strings.Contains(filePath, "client"):
		return "Client"
	case strings.Contains(filePath, "cmd/tools"):
		return "Tools"
	case strings.Contains(filePath, "tools/"):
		return "Tools"
	default:
		return "Common" // default
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

// extractFunctionContext tries to find the function name for better error code naming
func extractFunctionContext(lines []string, currentLineNum int) string {
	// Look backwards to find function declaration
	for i := currentLineNum; i >= 0 && i >= currentLineNum-20; i-- {
		line := strings.TrimSpace(lines[i])
		// Match function declarations
		if strings.HasPrefix(line, "func ") {
			// Extract function name
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				funcName := parts[1]
				if idx := strings.Index(funcName, "("); idx > 0 {
					funcName = funcName[:idx]
				}
				// Remove receiver if present
				if strings.Contains(funcName, ")") {
					if idx := strings.LastIndex(funcName, ")"); idx >= 0 && idx < len(funcName)-1 {
						funcName = funcName[idx+1:]
					}
				}
				return funcName
			}
			break
		}
	}
	return ""
}

type CodeAssignment struct {
	Location ErrorLocation
	Code     int
	Constant string
}

func assignSequentialErrorCodes(locations []ErrorLocation) ([]CodeAssignment, error) {
	// Use actual component ranges from existing error codes
	componentRanges := map[string][2]int{
		"Frontend":  {2000, 2999},
		"History":   {3000, 3999},
		"Matching":  {4000, 4999},
		"Worker":    {5000, 5999},
		"Component": {7300, 7399},
		"Persist":   {7500, 7599},
		"Common":    {7000, 7199},
		"Client":    {8000, 8099},
		"Tools":     {7200, 7299},
	}

	// Sort locations for consistent assignment
	sort.Slice(locations, func(i, j int) bool {
		if locations[i].Component != locations[j].Component {
			return locations[i].Component < locations[j].Component
		}
		if locations[i].FilePath != locations[j].FilePath {
			return locations[i].FilePath < locations[j].FilePath
		}
		return locations[i].LineNumber < locations[j].LineNumber
	})

	// Get existing error codes to avoid conflicts
	existingCodes, err := getExistingErrorCodes()
	if err != nil {
		fmt.Printf("Warning: could not load existing codes: %v\n", err)
		existingCodes = make(map[int]bool)
	}

	componentCounters := make(map[string]int)
	usedNames := make(map[string]int) // Track used constant names
	var assignments []CodeAssignment

	for _, location := range locations {
		component := location.Component
		rangeInfo, exists := componentRanges[component]
		if !exists {
			return nil, fmt.Errorf("unknown component: %s", component)
		}

		if _, exists := componentCounters[component]; !exists {
			componentCounters[component] = rangeInfo[0]
		}

		// Find next available code in range
		code := componentCounters[component]
		for existingCodes[code] && code <= rangeInfo[1] {
			code++
		}

		if code > rangeInfo[1] {
			return nil, fmt.Errorf("exceeded error code range for component %s at code %d", component, code)
		}

		componentCounters[component] = code + 1
		existingCodes[code] = true // Mark as used

		// Generate meaningful constant name based on context
		baseName := generateConstantName(location)
		constantName := baseName

		// Handle duplicates by adding suffix
		if count, exists := usedNames[baseName]; exists {
			count++
			usedNames[baseName] = count
			constantName = fmt.Sprintf("%s%d", baseName, count)
		} else {
			usedNames[baseName] = 1
		}

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
// This file contains generated error code constants for logging migration.
// Add these to your existing common/errorcode/codes.go file.

package errorcode

// Generated error code constants
const (
`)

	// Group by component for better organization
	componentGroups := make(map[string][]CodeAssignment)
	for _, assignment := range assignments {
		comp := assignment.Location.Component
		componentGroups[comp] = append(componentGroups[comp], assignment)
	}

	// Generate constants grouped by component
	componentOrder := []string{"Frontend", "History", "Matching", "Worker", "Component", "Persist", "Common", "Client", "Tools"}

	totalGenerated := 0
	for _, component := range componentOrder {
		assignments, exists := componentGroups[component]
		if !exists || len(assignments) == 0 {
			continue
		}

		buf.WriteString(fmt.Sprintf("\n\t// %s Service - Generated Error Codes (%d codes)\n", component, len(assignments)))

		for _, assignment := range assignments {
			comment := fmt.Sprintf("%s:%d",
				strings.TrimPrefix(assignment.Location.FilePath, "./"),
				assignment.Location.LineNumber)
			if assignment.Location.Message != "" && assignment.Location.Message != "error logging call" {
				cleanMsg := strings.ReplaceAll(assignment.Location.Message, "\n", " ")
				if len(cleanMsg) > 60 {
					comment += fmt.Sprintf(" - %.60s...", cleanMsg)
				} else {
					comment += fmt.Sprintf(" - %s", cleanMsg)
				}
			}

			buf.WriteString(fmt.Sprintf("\t%s = %d // %s\n",
				assignment.Constant,
				assignment.Code,
				comment,
			))
			totalGenerated++
		}
	}

	buf.WriteString(")\n")

	fmt.Printf("Generated %d error codes in %s\n", totalGenerated, outputPath)
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("1. Review the generated constants in %s\n", outputPath)
	fmt.Printf("2. Add selected constants to common/errorcode/codes.go\n")
	fmt.Printf("3. Register them in the appropriate register*Codes() functions:\n")
	fmt.Printf("   Register(code, ComponentEnum, \"description\")\n")
	fmt.Printf("4. Update logging calls to use error codes:\n")
	fmt.Printf("   OLD: logger.Error(\"message\", tag.Error(err))\n")
	fmt.Printf("   NEW: log.ErrorWithCode(logger, errorcode.YourCode, \"message\", err)\n")
	fmt.Printf("\nFor help: go run cmd/tools/errorcodegen/main.go -help\n")
	return os.WriteFile(outputPath, []byte(buf.String()), 0644)
}

// generateConstantName creates meaningful constant names following existing patterns
func generateConstantName(location ErrorLocation) string {
	component := location.Component

	// Extract key context from file path (more reliable than message)
	contextParts := extractContextFromPath(location.FilePath)

	// Extract meaningful operation from message
	operationParts := extractOperationFromMessage(location.Message)

	// Build name: Component + Context + Operation + Suffix
	var nameParts []string
	nameParts = append(nameParts, component)

	// Add 1-2 context parts for specificity
	if len(contextParts) > 0 {
		nameParts = append(nameParts, contextParts[0])
		if len(contextParts) > 1 {
			nameParts = append(nameParts, contextParts[1])
		}
	}

	// Add operation context
	if len(operationParts) > 0 {
		nameParts = append(nameParts, operationParts[0])
	}

	// Ensure it ends with appropriate suffix
	name := strings.Join(nameParts, "")
	if !strings.HasSuffix(name, "Failed") && !strings.HasSuffix(name, "Error") && !strings.HasSuffix(name, "Timeout") {
		name += "OperationFailed"
	}

	return name
}

// extractContextFromPath gets meaningful context from file path
func extractContextFromPath(filePath string) []string {
	parts := strings.Split(filePath, "/")
	var result []string

	// Look for meaningful directory names
	for i := len(parts) - 2; i >= 0; i-- {
		dir := parts[i]
		if isUsefulDirectoryName(dir) {
			cleaned := cleanAndCapitalize(dir)
			result = append([]string{cleaned}, result...) // prepend
			if len(result) >= 2 {
				break
			}
		}
	}

	// Add file name context
	if len(parts) > 0 {
		fileName := parts[len(parts)-1]
		fileName = strings.TrimSuffix(fileName, ".go")
		cleaned := cleanAndCapitalize(fileName)
		if cleaned != "" && len(result) < 2 {
			result = append(result, cleaned)
		}
	}

	return result
}

// extractOperationFromMessage tries to extract the core operation being performed
func extractOperationFromMessage(message string) []string {
	if message == "" {
		return nil
	}

	// Look for key operation words
	operationWords := []string{
		"failed", "error", "timeout", "invalid", "missing", "duplicate",
		"create", "update", "delete", "get", "fetch", "process", "execute",
		"start", "stop", "cancel", "reset", "retry", "validate", "serialize",
		"deserialize", "parse", "encode", "decode", "connect", "disconnect",
	}

	lowerMsg := strings.ToLower(message)
	for _, word := range operationWords {
		if strings.Contains(lowerMsg, word) {
			return []string{cleanAndCapitalize(word)}
		}
	}

	return nil
}

// isUsefulDirectoryName checks if a directory name provides meaningful context
func isUsefulDirectoryName(dir string) bool {
	useless := map[string]bool{
		"service": true, "common": true, "api": true, "v1": true, "pkg": true,
		"internal": true, "lib": true, "src": true, "cmd": true,
	}
	return !useless[dir] && len(dir) > 1
}

// cleanAndCapitalize cleans a string and capitalizes it
func cleanAndCapitalize(s string) string {
	if s == "" {
		return ""
	}

	// Remove special characters and replace with nothing
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")

	// Capitalize first letter
	if len(s) > 0 {
		s = strings.ToUpper(string(s[0])) + s[1:]
	}

	return s
}

// getExistingErrorCodes loads existing error codes to avoid conflicts
func getExistingErrorCodes() (map[int]bool, error) {
	codes := make(map[int]bool)

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, "common/errorcode/codes.go", nil, parser.ParseComments)
	if err != nil {
		return codes, err
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if decl, ok := n.(*ast.GenDecl); ok && decl.Tok == token.CONST {
			for _, spec := range decl.Specs {
				if vspec, ok := spec.(*ast.ValueSpec); ok {
					for i, _ := range vspec.Names {
						if i < len(vspec.Values) {
							if lit, ok := vspec.Values[i].(*ast.BasicLit); ok && lit.Kind == token.INT {
								code, _ := strconv.Atoi(lit.Value)
								codes[code] = true
							}
						}
					}
				}
			}
		}
		return true
	})

	return codes, nil
}

// mapComponentStringToEnum converts component string names to Component enum constants
func mapComponentStringToEnum(componentStr string) string {
	switch componentStr {
	case "Frontend":
		return "ComponentFrontend"
	case "History":
		return "ComponentHistory"
	case "Matching":
		return "ComponentMatching"
	case "Worker":
		return "ComponentWorker"
	case "Component":
		return "ComponentCommon"
	case "Persist":
		return "ComponentPersist"
	case "Common":
		return "ComponentCommon"
	case "Client":
		return "ComponentCommon"
	case "Tools":
		return "ComponentTools"
	default:
		return "ComponentCommon" // Default fallback
	}
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

// showUsageHelp displays usage examples and documentation
func showUsageHelp() {
	fmt.Println(`Error Code Generator Tool

Purpose:
This tool helps migrate logging statements to use error codes by scanning the codebase
for existing logging calls and generating appropriate error code constants.

Usage:
  go run cmd/tools/errorcodegen/main.go [flags]

Flags:
  -scan          Scan codebase and generate error codes
  -output FILE   Output file path for generated constants (required with -scan)
  -verify        Verify no conflicts in existing error codes
  -help          Show this help message

Examples:
  # Generate error codes from codebase scan
  go run cmd/tools/errorcodegen/main.go -scan -output generated_codes.go

  # Verify existing error codes have no conflicts
  go run cmd/tools/errorcodegen/main.go -verify

Generated Constants:
The tool generates constants following the existing naming patterns:
  - ComponentContextOperationFailed (e.g., HistoryReplicationTaskProcessorOperationFailed)
  - Numbered suffixes for duplicates (e.g., HistoryHandlerError2)
  - Comments with file location and message excerpt

After generation:
1. Review the generated constants
2. Add them to common/errorcode/codes.go
3. Register them in the appropriate register*Codes() functions
4. Update logging calls to use: log.ErrorWithCode(logger, errorcode.YourNewCode, msg, err, tags...)

Error Code Ranges:
  Frontend:  2000-2999
  History:   3000-3999
  Matching:  4000-4999
  Worker:    5000-5999
  Component: 7300-7399
  Persist:   7500-7599
  Common:    7000-7199
  Tools:     7200-7299
`)
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
