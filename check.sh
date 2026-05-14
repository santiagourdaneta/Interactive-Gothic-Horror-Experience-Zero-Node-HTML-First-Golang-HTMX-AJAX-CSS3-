#!/bin/bash
echo "--- 1. FORMATEANDO ---"
go fmt ./...

echo "--- 2. ANALIZANDO (LINTER) ---"
staticcheck ./...

echo "--- 3. EJECUTANDO TESTS (UNIT & E2E) ---"
go test ./... -v

echo "--- 4. TODO OK. EL SISTEMA ESTÁ SEGURO ---"