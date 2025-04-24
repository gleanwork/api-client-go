#!/bin/bash

# Install the speakeasy CLI
curl -fsSL https://raw.githubusercontent.com/speakeasy-api/speakeasy/main/install.sh | sh

# Setup samples directory
rmdir samples || true
mkdir samples

# Go module commands
go mod download
go mod tidy

# Generate starter usage sample with speakeasy
speakeasy generate usage -s registry.speakeasyapi.dev/glean-el2/sdk/glean-api-specs -l go -o samples/root.go