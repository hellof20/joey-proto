#!/bin/bash

# Joey Proto Repository Deployment Script
# This script initializes the git repository and pushes to GitHub

set -e

echo "========================================="
echo "Joey Proto Repository Setup"
echo "========================================="
echo ""

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "Error: go.mod not found. Please run this script from joey-proto directory"
    exit 1
fi

# Initialize git if not already initialized
if [ ! -d ".git" ]; then
    echo "Initializing git repository..."
    git init
    echo "✓ Git initialized"
else
    echo "✓ Git repository already exists"
fi

# Add all files
echo ""
echo "Adding files to git..."
git add .
echo "✓ Files added"

# Create initial commit
echo ""
echo "Creating initial commit..."
git commit -m "Initial commit: Joey Proto v0.1.0

- Add gRPC protocol definitions (game_agent.proto)
- Add shared logger utilities
- Setup Go module (github.com/hellof20/joey-proto)
- Add Makefile for proto generation
- Add GitHub Actions workflow
- Add documentation"
echo "✓ Initial commit created"

# Add remote (will fail if already exists, but that's OK)
echo ""
echo "Adding remote origin..."
git remote add origin https://github.com/hellof20/joey-proto.git 2>/dev/null || echo "✓ Remote already exists"

# Create and switch to main branch
echo ""
echo "Setting up main branch..."
git branch -M main
echo "✓ Main branch ready"

# Push to GitHub
echo ""
echo "Pushing to GitHub..."
echo "You may be prompted for your GitHub credentials..."
git push -u origin main

echo ""
echo "========================================="
echo "✓ Deployment Complete!"
echo "========================================="
echo ""
echo "Next steps:"
echo "1. Visit: https://github.com/hellof20/joey-proto"
echo "2. Generate proto files locally: make proto"
echo "3. Create v0.1.0 tag:"
echo "   git tag -a v0.1.0 -m 'Release v0.1.0'"
echo "   git push origin v0.1.0"
echo ""
