# Error Code Generator

This tool automatically scans the Temporal codebase for error logging patterns and generates unique error codes for each location.

## Features

- **Automated Scanning**: Finds all error logging calls (`.Error()`, `.Warn()`, `.Fatal()`, etc.)
- **Component Classification**: Automatically assigns components based on file paths
- **Sequential Assignment**: Assigns codes sequentially within component ranges
- **Validation**: Prevents duplicate codes and range violations
- **Build Integration**: Can be integrated into the build process

## Usage

### Generate Error Codes

Scan the entire codebase and generate error codes:

```bash
go run cmd/tools/errorcodegen/main.go -scan -output common/errorcode/generated_codes.go
```

Or use the convenient script:

```bash
./cmd/tools/errorcodegen/generate.sh
```

### Verify Error Codes

Check existing error codes for conflicts:

```bash
go run cmd/tools/errorcodegen/main.go -verify
```

### Manual Generation (Legacy)

```bash
go run cmd/tools/errorcodegen/main.go -output output_file.go
```

## Component Ranges

| Component | Range | Description |
|-----------|--------|-------------|
| INFRA | 1000-1999 | Infrastructure (persistence, clustering) |
| FRONT | 2000-2999 | Frontend service |
| HIST | 3000-3999 | History service |
| MATCH | 4000-4999 | Matching service |
| WORK | 5000-5999 | Worker service |
| WF | 6000-6999 | Workflow engine |
| SYS | 7000-7999 | System utilities |
| EXT | 8000-8999 | External integrations |

## Generated Code Format

The tool generates constants in the format `COMPONENT_CODE`:

```go
const (
    // INFRA Service Error Codes (493 codes)
    INFRA_1000 = 1000 // common/persistence/cassandra/factory.go:44 - unable to initialize cassandra session
    INFRA_1001 = 1001 // common/persistence/cassandra/helpers.go:19 - drop keyspace error
    
    // FRONT Service Error Codes (119 codes)  
    FRONT_2000 = 2000 // service/frontend/admin_handler.go:45 - error logging call
    FRONT_2001 = 2001 // service/frontend/admin_handler.go:67 - failed to get namespace
)
```

And automatic registration functions:

```go
func init() {
    registerINFRACodes()
    registerFRONTCodes()
    // ... other components
}

func registerINFRACodes() {
    Register(INFRA_1000, "INFRA", "unable to initialize cassandra session")
    Register(INFRA_1001, "INFRA", "drop keyspace error")
    // ... other codes
}
```

## Scanning Process

1. **File Discovery**: Walks through source directories (`service/`, `common/`, `client/`)
2. **Pattern Matching**: Uses regex to find error logging calls
3. **Message Extraction**: Extracts error messages from string literals  
4. **Component Classification**: Assigns components based on file paths
5. **Sequential Assignment**: Assigns codes in order within component ranges
6. **Code Generation**: Outputs Go constants and registration functions

## Statistics

Recent scan results:
- **Total Error Locations**: 1,436
- **INFRA**: 493 codes (persistence layer)
- **HIST**: 503 codes (workflow execution)
- **WORK**: 217 codes (background workers)
- **FRONT**: 119 codes (API layer)
- **MATCH**: 67 codes (task queues)
- **SYS**: 37 codes (utilities)

## Integration with Build

The error code generation can be integrated into the build process:

1. Add to `Makefile`:
```makefile
.PHONY: generate-error-codes
generate-error-codes:
	./cmd/tools/errorcodegen/generate.sh

.PHONY: verify-error-codes  
verify-error-codes:
	go run cmd/tools/errorcodegen/main.go -verify
```

2. Include in `go-generate`:
```makefile
go-generate: generate-error-codes
	go generate ./...
```

## Testing

Test the generated codes:

```bash
go run cmd/tools/errorcodegen/test_generated.go
```

Run validation:

```bash
go test ./common/errorcode/ -run TestErrorCodeRegistration
```