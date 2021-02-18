.PHONY: help
help: ## - Show this help message
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build-common
build-common: ## - execute build common tasks clean and mod tidy
	@ go version
	@ go clean
	@ go mod tidy

.PHONY: build-release
build-release: ## - build a release linux elf(binary)
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -a -installsuffix cgo -o "bin/release/clientapi-restapi-server" main.go

.PHONY: build-static-release
build-static-release: ## - build a static release linux elf(binary)
	@ CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o "bin/release/clientapi-restapi-server" main.go

.PHONY: build-debug
build-debug: build-common ## - build a debug binary to the current platform (windows, linux or darwin(mac))
	@ echo cleaning...
	@ rm -f bin/debug/clientapi-restapi-server
	@ echo building...
	@ go build -o "bin/debug/clientapi-restapi-server" main.go
	@ ls -lah bin/debug/clientapi-restapi-server

.PHONY: test
test: build-common ## - execute go test command
	@go test -v -cover ./...

.PHONY: ci-lint
ci-lint: ## - runs golangci-lint
	@ golangci-lint run --enable-all -D forbidigo -D gochecknoglobals -D gofumpt -D gofmt -D nlreturn

docker-build: ## - build docker image
	@ docker build -f Dockerfile -t fafg/clientapi:latest .

.PHONY: docker-scan
docker-scan: ## - Scan for known vulnerabilities
	@printf "\033[32m\xE2\x9c\x93 Scan for known vulnerabilities the smallest and secured golang docker image based on scratch\n\033[0m"
	@docker scan -f Dockerfile clientapi

.PHONY: sonar-start
sonar-start: ## - start sonar qube locally with docker (you will need docker installed in your machine)
	# after start, setup a new project with the name sms-local and a new token sms-token, fill the token against the -Dsonar.login= parameter.
	# login with user: admin pwd: admin
	@$(SHELL) ../scripts/sonar-start.sh

.PHONY: sonar-stop
sonar-stop: ## - stop sonar qube docker container
	@docker stop sonarqube

.PHONY: sonar-scan
sonar-scan: build-debug ## - runs build and then sonar scanner (make sure you have installed sonar-scanner and you have it in your path)
	sonar-scanner -Dsonar.projectKey="clientapi" \
				  -Dsonar.projectName="clientapi" \
				  -Dsonar.sources=. \
				  -Dsonar.exclusions=internal/** \
				  -Dsonar.host.url="http://localhost:9000" \
				  -Dsonar.login="<token here>"
