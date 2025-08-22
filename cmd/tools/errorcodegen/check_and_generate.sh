#!/bin/bash
# The MIT License
#
# Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in
# all copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
# THE SOFTWARE.

set -euo pipefail

# Change to the root directory of the project
cd "$(dirname "$0")/../../.."

GENERATED_FILE="common/errorcode/generated_codes.go"
CHECKSUM_FILE="$GENERATED_FILE.checksum"

# Function to calculate checksum of all source files that could contain error logging
calculate_source_checksum() {
    find service/ common/ client/ -name "*.go" -not -name "*_test.go" -not -path "*/generated_codes.go" 2>/dev/null | \
        sort | \
        xargs sha256sum 2>/dev/null | \
        sha256sum | \
        cut -d' ' -f1
}

# Check if we need to regenerate
should_regenerate=false

if [[ ! -f "$GENERATED_FILE" ]]; then
    echo "Generated file $GENERATED_FILE does not exist, generating..."
    should_regenerate=true
elif [[ ! -f "$CHECKSUM_FILE" ]]; then
    echo "Checksum file $CHECKSUM_FILE does not exist, regenerating..."
    should_regenerate=true
else
    current_checksum=$(calculate_source_checksum)
    stored_checksum=$(cat "$CHECKSUM_FILE" 2>/dev/null || echo "")
    
    if [[ "$current_checksum" != "$stored_checksum" ]]; then
        echo "Source files have changed, regenerating error codes..."
        should_regenerate=true
    else
        echo "Source files unchanged, skipping error code generation."
    fi
fi

if [[ "$should_regenerate" == "true" ]]; then
    echo "Generating error codes..."
    
    # Generate error codes
    go run cmd/tools/errorcodegen/main.go -scan -output "$GENERATED_FILE"
    
    # Store the new checksum
    current_checksum=$(calculate_source_checksum)
    echo "$current_checksum" > "$CHECKSUM_FILE"
    
    echo "Error code generation complete."
    
    # Verify the generated codes
    echo "Verifying generated error codes..."
    go run cmd/tools/errorcodegen/main.go -verify
    echo "Verification passed."
else
    # Still verify existing codes
    echo "Verifying existing error codes..."
    go run cmd/tools/errorcodegen/main.go -verify
    echo "Verification passed."
fi