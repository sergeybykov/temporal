package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/server/common/errorcode"
)

func TestErrorCodeGeneration(t *testing.T) {
	// Test that error codes are generated and registered correctly

	// Get all registered codes
	allCodes := errorcode.GetAllCodes()
	require.Greater(t, len(allCodes), 200, "Should have generated many error codes")

	// Test component distribution
	componentCounts := make(map[string]int)
	for _, code := range allCodes {
		componentCounts[code.Component.String()]++
	}

	// Verify all expected components exist
	expectedComponents := []string{"INFRA", "FRONT", "HIST", "MATCH", "WORK"}
	for _, comp := range expectedComponents {
		require.Greater(t, componentCounts[comp], 0, "Component %s should have codes", comp)
	}

	// Test that some components have reasonable numbers of codes
	require.Greater(t, componentCounts["INFRA"], 5, "INFRA should have codes")
	require.Greater(t, componentCounts["HIST"], 5, "HIST should have codes")
}

func TestErrorCodeRanges(t *testing.T) {
	// Test that all codes are in correct ranges
	allCodes := errorcode.GetAllCodes()

	componentRanges := map[string][2]int{
		"INFRA":    {1000, 1999},
		"FRONT":    {2000, 2999},
		"HIST":     {3000, 3999},
		"MATCH":    {4000, 4999},
		"WORK":     {5000, 5999},
		"WF":       {6000, 6999},
		"SYS":      {7000, 7999},
		"EXT":      {8000, 8999},
		"RES":      {9000, 9999},
		"COMMON":   {7000, 9099},
		"PERSIST":  {7500, 9099},
		"TOOLS":    {7200, 7299},
		"SCHEMA":   {8200, 8299},
		"ARCHIVER": {7100, 9099},
		"TEST":     {7600, 7699},
	}

	for _, code := range allCodes {
		componentStr := code.Component.String()
		r, exists := componentRanges[componentStr]
		require.True(t, exists, "Component %s should have a defined range", componentStr)
		require.GreaterOrEqual(t, code.Code, r[0], "Code %d should be >= %d for component %s", code.Code, r[0], componentStr)
		require.LessOrEqual(t, code.Code, r[1], "Code %d should be <= %d for component %s", code.Code, r[1], componentStr)
	}
}

func TestErrorCodeUniqueness(t *testing.T) {
	// Test that all error codes are unique
	allCodes := errorcode.GetAllCodes()
	seenCodes := make(map[int]bool)

	for _, code := range allCodes {
		require.False(t, seenCodes[code.Code], "Error code %d should be unique", code.Code)
		seenCodes[code.Code] = true
	}
}

func TestSpecificGeneratedCodes(t *testing.T) {
	// Test some specific generated codes to ensure they work
	testCodes := []struct {
		code      int
		component string
	}{
		{1101, "INFRA"}, // Known INFRA code
		{2001, "FRONT"}, // Known FRONT code
		{3001, "HIST"},  // Known HIST code
		{4001, "MATCH"}, // Known MATCH code
		{5001, "WORK"},  // Known WORK code
		{7001, "COMMON"}, // Known COMMON code
	}

	for _, tc := range testCodes {
		ec, exists := errorcode.Get(tc.code)
		require.True(t, exists, "Error code %d should exist", tc.code)
		require.Equal(t, tc.component, ec.Component.String(), "Error code %d should belong to component %s", tc.code, tc.component)
		require.NotEmpty(t, ec.Message, "Error code %d should have a message", tc.code)
	}
}
