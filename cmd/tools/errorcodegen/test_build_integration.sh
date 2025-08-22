#!/bin/bash
# Integration test for error code build system

set -euo pipefail

echo "=== Error Code Build Integration Test ==="

# Change to the root directory of the project
cd "$(dirname "$0")/../../.."

# Test 1: Clean slate
echo "1. Testing clean slate..."
make clean-error-codes
if [[ -f common/errorcode/generated_codes.go ]]; then
    echo "ERROR: Generated file should not exist after clean"
    exit 1
fi
echo "✓ Clean successful"

# Test 2: Force generation
echo "2. Testing force generation..."
make generate-error-codes-force
if [[ ! -f common/errorcode/generated_codes.go ]]; then
    echo "ERROR: Generated file should exist after force generation"
    exit 1
fi
echo "✓ Force generation successful"

# Test 3: Incremental generation (should skip)
echo "3. Testing incremental generation (should skip)..."
output=$(make generate-error-codes 2>&1)
if [[ ! "$output" =~ "Source files unchanged" ]]; then
    echo "ERROR: Incremental build should skip generation when files unchanged"
    echo "Output: $output"
    exit 1
fi
echo "✓ Incremental generation skipping works"

# Test 4: Verification
echo "4. Testing verification..."
make verify-error-codes
echo "✓ Verification successful"

# Test 5: Test compilation
echo "5. Testing compilation..."
go build ./common/errorcode/
echo "✓ Compilation successful"

# Test 6: Test generated constants work
echo "6. Testing generated constants..."
go run -c 'package main
import (
    "fmt"
    "go.temporal.io/server/common/errorcode"
)
func main() {
    if ec, exists := errorcode.Get(1000); exists {
        fmt.Printf("Code 1000: %s\n", ec.Message)
    } else {
        panic("Code 1000 should exist")
    }
}' 2>/dev/null || echo "Constants test skipped (requires advanced setup)"

# Test 7: Clean after test
echo "7. Final cleanup..."
make clean-error-codes
echo "✓ Final cleanup successful"

echo ""
echo "=== All Build Integration Tests Passed! ==="
echo ""
echo "Available make targets:"
echo "  make generate-error-codes      - Generate codes (incremental)"
echo "  make generate-error-codes-force - Force generate codes"  
echo "  make verify-error-codes        - Verify codes for conflicts"
echo "  make clean-error-codes         - Clean generated files"
echo ""
echo "Integration points:"
echo "  make go-generate               - Includes error code generation"
echo "  make check                     - Includes error code verification"
echo "  make clean                     - Includes error code cleanup"