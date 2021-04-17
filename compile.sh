#!/usr/bin/env bash
set -e
PROTOC_VERSION=$(protoc --version | awk '{print $2}')
if [ "$PROTOC_VERSION" != "3.7.1" ]; then
    echo "protoc version should be 3.7.1"
    exit 1
fi

## golang
rm -rf go && mkdir -p go
#protoc -Iproto --go_opt=paths=source_relative --go_out=go proto/*.proto

services="gateway common"
for service in $services; do
    # golang
    rm -rf go/$service && mkdir -p go/$service
    find proto/$service -type f -name '*.proto' -exec protoc -Iproto --go_opt=paths=source_relative --go_out=plugins=grpc:go {} \;
done

# grpc golang
rm -rf go/grpc && mkdir -p go/grpc
for service in center gateway chat validator; do 
    protoc -Iproto --go_opt=paths=source_relative --go_out=plugins=grpc:go proto/grpc/$service/$service.proto
done

# for service in center gateway chat; do
#     protoc -Iproto -I/usr/local/include --go_opt=paths=source_relative --govalidators_out=go --go_out=plugins=grpc:go proto/grpc/$service/$service.proto
# done
