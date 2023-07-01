#!/bin/sh -e

INPUT_DIR="openapi"
CONFIG_DIR="openapi/config"
OUTPUT_DIR="interface/web"
oapi-codegen --config ${CONFIG_DIR}/types_gen.yaml ${INPUT_DIR}/petstore-expanded.yaml >| ${OUTPUT_DIR}/types_gen.go
oapi-codegen -generate server -package web ${INPUT_DIR}/petstore-expanded.yaml >| ${OUTPUT_DIR}/server_gen.go
oapi-codegen -generate spec -package web ${INPUT_DIR}/petstore-expanded.yaml >| ${OUTPUT_DIR}/spec_gen.go
