#!/usr/bin/env bash
set -o errexit

go build -o scripts/hcl-parser scripts/hcl-parser.go
scripts/hcl-parser aws
scripts/hcl-parser azurerm
scripts/hcl-parser gcp
rm -f scripts/hcl-parser
go fmt template/*

echo "----------------------------------------"
echo "--- !!! updated template structs !!! ---"
echo "----------------------------------------"
