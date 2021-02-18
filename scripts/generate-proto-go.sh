#!/usr/bin/env bash

protoc -I "${PWD}/../portservice" --go_out="${PWD}/../" "${PWD}/../portservice/internal/proto-files/domain/port.proto"
protoc -I "${PWD}/../" --go_out=plugins=grpc:"${PWD}/../" "${PWD}/../portservice/internal/proto-files/service/port-service.proto"

protoc -I "${PWD}/../portservice" --go_out="${PWD}/../clientapi/" "${PWD}/../portservice/internal/proto-files/domain/port.proto"
protoc -I "${PWD}/../" --go_out=plugins=grpc:"${PWD}/../clientapi/" "${PWD}/../portservice/internal/proto-files/service/port-service.proto"