# Build Status
![Go-Docker-Example](https://github.com/fafg/go-docker-example/workflows/Go-Docker-Example/badge.svg)

# solution
The solution is composed of two service plus a database used to store and search data.

Three docker containers are necessary to run locally:
1. client api
2. port service (serving a grpc service)
3. mongodb

what i have used so far:
> grpc, mongodb with full text search, echo lib to rest api, github actions to ci, docker hub and docker containers, make file, a couple of shell scripts,ddd. 

# how to run

Before you run, please change the docker-compose.yaml file appropriately with the right folder path to the json file.

```yaml
#client api service
volumes:
  #do not add the file name here, only the path
  - /path/to/your/json/file/:/go/file/
```

> just execute ```run.sh```

Check the logs
```shell
docker logs <clientapi-container-name> -f
docker logs <portservice-container-name> -f
```

over each project folder you can execute the command to see what is available, like this:
```shell
make
✓ usage: make [target]

build-common                   - execute build common tasks clean and mod tidy
build-debug                    - build a debug binary to the current platform (windows, linux or darwin(mac))
build-release                  - build a release linux elf(binary)
build-static-release           - build a static release linux elf(binary)
ci-lint                        - runs golangci-lint
docker-build                   - build docker image
docker-scan                    - Scan for known vulnerabilities
generate-go-proto              - generate golang proto files
help                           - Show this help message
sonar-scan                     - runs build and then sonar scanner (make sure you have installed sonar-scanner and you have it in your path)
sonar-start                    - start sonar qube locally with docker (you will need docker installed in your machine)
sonar-stop                     - stop sonar qube docker container
test                           - execute go test command

```


# What can be improved
> configuration, failure recovery, unit test, add more features, usage of k3d for kubernetes local deployment.