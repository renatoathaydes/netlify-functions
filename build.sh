#!/bin/bash

set -e

FUNCTIONS=website/netlify/functions/analytics

echo "Build go"
cd $FUNCTIONS && go test .
echo "Built successfully"
