package main

import (
    "context"

    log "github.com/sirupsen/logrus"
    "google.golang.org/grpc"

    "portservice/internal/grpc/domain"
    "portservice/internal/grpc/service"
    "portservice/logger"
)

func main() {
    logger.Init()
    serverAddress := "localhost:7070"

    conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

    if e != nil {
        panic(e)
    }
    defer func() {
        err := conn.Close()
        if err != nil {
            log.Errorf("error on closing connection: %v\n", err)
        }
    }()

    client := service.NewPortDomainServiceClient(conn)

    for range [10]int{} {
        domainModel := domain.Airport{
            Codename: "AAA",
            City: "Pliezhausen",
        }

        if responseMessage, err := client.Insert(context.Background(), &domainModel); err != nil {
            log.Errorf("Error: %v\n", err)
        } else {
            log.Infof("Record Inserted: %v\n", responseMessage)
        }
    }
}