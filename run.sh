#!/usr/bin/env bash

docker-compose up -d

echo "sleeping 3s to make sure container is up."
sleep 3s

echo "running functional test"
./scripts/functional-test.sh
