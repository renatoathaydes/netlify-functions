#!/bin/bash

set -e

echo "Build go"
go build .
echo "Tests"
go test .
echo "Clean website/netlify dir"
rm -rf website/netifly/
echo "Copy go files to website/netlify dir"
mkdir -p website/netlify/functions/analytics/
cp *.go website/netlify/functions/analytics/
echo "Built successfully"