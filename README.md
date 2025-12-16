# Joey Proto

Joey gRPC Protocol Definitions and Shared Utilities

## Overview

This repository contains the shared protocol definitions and utilities used by both the Joey server and client applications. It serves as the contract layer between the server and client, ensuring consistent communication.

## Contents

- **proto/** - gRPC Protocol Buffer definitions
  - `game_agent.proto` - Core service definitions for client-server communication

- **logger/** - Shared logging utilities
  - Structured logging based on logrus
  - Context-aware logging helpers
  - Support for both JSON and text formats

## Installation

```bash
go get github.com/hellof20/joey-proto
```

## Usage

### In Server Code

```go
import (
    "github.com/hellof20/joey-proto/proto"
    "github.com/hellof20/joey-proto/logger"
)
```

### In Client Code

```go
import (
    "github.com/hellof20/joey-proto/proto"
    "github.com/hellof20/joey-proto/logger"
)
```

## Development

### Prerequisites

- Go 1.24.3 or higher
- Protocol Buffers compiler (`protoc`)
- protoc-gen-go and protoc-gen-go-grpc plugins

Install the protoc plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

### Generating Go Code from Proto

```bash
make proto
```

This will generate:
- `proto/game_agent.pb.go` - Protocol Buffer message definitions
- `proto/game_agent_grpc.pb.go` - gRPC service definitions

### Clean Generated Files

```bash
make clean
```

## Protocol Version Compatibility

This library follows semantic versioning:

- **Major version** (v1.x.x → v2.x.x): Breaking changes to the protocol
- **Minor version** (v1.0.x → v1.1.x): Backward-compatible additions
- **Patch version** (v1.0.0 → v1.0.1): Bug fixes and documentation

When adding new fields to proto messages, always use new field numbers to maintain backward compatibility.

## Dependencies

- `github.com/gin-gonic/gin` - HTTP framework (for logger context helpers)
- `github.com/sirupsen/logrus` - Structured logging
- `google.golang.org/grpc` - gRPC framework
- `google.golang.org/protobuf` - Protocol Buffers

## License

MIT License

## Related Repositories

- [joey-server](https://github.com/hellof20/joey-server) - Joey server application
- [joey-client](https://github.com/hellof20/joey-client) - Joey client application
