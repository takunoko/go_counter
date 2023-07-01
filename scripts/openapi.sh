#!/bin/sh -e

INPUT_DIR="openapi"
OUTPUT_DIR="interface/web"
## oapi-codegen -generate types,server -package web ${INPUT_DIR}/petstore-expanded.yaml > ${OUTPUT_DIR}/web_gen.go
oapi-codegen -package web ${INPUT_DIR}/petstore-expanded.yaml > ${OUTPUT_DIR}/web_gen.go
