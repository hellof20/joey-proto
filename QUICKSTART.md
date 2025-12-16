# Joey Proto - Quick Start Guide

## Step 1: Navigate to Setup Directory

```bash
cd /tmp/joey-proto-setup
```

## Step 2: Install Dependencies

```bash
go mod tidy
```

## Step 3: Generate Proto Files

```bash
make proto
```

This will create:
- `proto/game_agent.pb.go`
- `proto/game_agent_grpc.pb.go`

## Step 4: Deploy to GitHub

```bash
./DEPLOY.sh
```

This script will:
1. Initialize git repository
2. Add all files
3. Create initial commit
4. Add GitHub remote
5. Push to `main` branch

## Step 5: Create v0.1.0 Release Tag

After successful push, create the first version tag:

```bash
git tag -a v0.1.0 -m "Release v0.1.0 - Initial Joey Proto Library"
git push origin v0.1.0
```

## Verification

Visit your repository:
https://github.com/hellof20/joey-proto

You should see:
- ✓ All files committed
- ✓ Generated proto files
- ✓ Tag v0.1.0 available

## Using in Other Projects

Once published, you can use it in joey-server and joey-client:

```bash
go get github.com/hellof20/joey-proto@v0.1.0
```

Then in your Go code:

```go
import (
    "github.com/hellof20/joey-proto/proto"
    "github.com/hellof20/joey-proto/logger"
)
```

## Troubleshooting

### Permission Denied
If you get a permission error when running DEPLOY.sh:
```bash
chmod +x DEPLOY.sh
```

### Authentication Failed
Make sure you have:
1. GitHub Personal Access Token configured
2. Or SSH key set up

For HTTPS authentication:
```bash
git config --global credential.helper osxkeychain
```

### Proto Generation Fails
Make sure protoc is installed:
```bash
brew install protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Next Steps

After joey-proto is deployed:
1. Create joey-server repository
2. Create joey-client repository
3. Update import paths in both repos
