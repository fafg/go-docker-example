#!/usr/bin/env bash

docker-compose kill
docker rm -f go-docker-example_clientapi_1 go-docker-example_portservice_1 go-docker-example_mongo_1