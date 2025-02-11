#!/usr/bin/env bash

# Code generated by internal/generate/teamcity/provider_tests.go; DO NOT EDIT.

set -euo pipefail

go test \
    ./internal/acctest/... \
    ./internal/attrmap/... \
    ./internal/conns/... \
    ./internal/create/... \
    ./internal/enum/... \
    ./internal/envvar/... \
    ./internal/errs/... \
    ./internal/experimental/... \
    ./internal/flex/... \
    ./internal/framework/... \
    ./internal/generate/... \
    ./internal/logging/... \
    ./internal/maps/... \
    ./internal/provider/... \
    ./internal/sdktypes/... \
    ./internal/slices/... \
    ./internal/sweep/... \
    ./internal/tags/... \
    ./internal/tfresource/... \
    ./internal/types/... \
    ./internal/ujson/... \
    ./internal/vault/... \
    ./internal/verify/... \
    -json
