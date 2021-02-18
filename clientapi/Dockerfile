############################
# STEP 1 build
############################
FROM golang:alpine as builder

# install make and ssl certs
RUN apk update && apk upgrade && apk add --no-cache make ca-certificates

# set working dir into $GOPATH
WORKDIR $GOPATH/src/clientapi

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "10001" \
    "clientapi"

# use modules
COPY go.mod .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

# Build the static binary
RUN make build-static-release

############################
# STEP 2 build a small image
############################
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /go/src/clientapi/bin/release/clientapi-restapi-server /go/bin/clientapi-restapi-server

# Use an unprivileged user.
USER clientapi:clientapi

# Run the hello binary.
ENTRYPOINT ["/go/bin/clientapi-restapi-server"]