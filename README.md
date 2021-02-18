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
# What can be improved
> configuration, failure recovery, unit test, add more features, usage of k3d for kubernetes local deployment.