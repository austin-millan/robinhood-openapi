#!/bin/sh
# This script will codegen a go package for a server stub or client.

rm -rf openapi/

docker run \
    --rm \
    --user $(id -u):$(id -g) \
    -v ${PWD}:/local \
    openapitools/openapi-generator-cli generate \
        -i /local/swagger.yml \
        -g go \
        -o /local/openapi
rm openapi/api_*.go \
    openapi/.travis.yml \
    openapi/client.go \
    openapi/configuration.go \
    openapi/git_push.sh \
    openapi/go.mod \
    openapi/go.sum \
    openapi/response.go \
    openapi/api/openapi.yaml \
    openapi/docs/*Api.md

echo "module gitlab.com/austin-millan/robinhood-openapi/openapi
go 1.12" > go.mod
