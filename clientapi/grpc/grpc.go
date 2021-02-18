package grpc

import (
    "context"
    "os"

    log "github.com/sirupsen/logrus"
    "google.golang.org/grpc"

    "clientapi/internal/grpc/domain"
    "clientapi/internal/grpc/service"
)

var _connection *grpc.ClientConn
var _client service.PortDomainServiceClient

func Init() {
    var err error
    serverAddress :=  os.Getenv("PORT_SERVICE_ENDPOINT") //"localhost:7070"

    _connection, err = grpc.Dial(serverAddress, grpc.WithInsecure())
    if err != nil {
        log.Errorf("grpc connection could not be created: %v\n", err)
    }

    _client = service.NewPortDomainServiceClient(_connection)
    log.Info("grpc client has been initialized.")
}

func Close() {
    err := _connection.Close()
    if err != nil {
        log.Errorf("error on closing connection: %v\n", err)
    }
    log.Warn("grpc connection closed")
}

func SendAirportToRegistration(airport *domain.Airport) *service.InsertAirportResponse {
    if _client != nil {
        responseMessage, err := _client.Insert(context.Background(), airport)
        if err != nil {
            log.Errorf("Error: %v\n", err)
        } else {
            log.Infof("Record Inserted: %v\n", responseMessage)
        }

        return responseMessage
    }

    return nil
}

func SearchAirportByName(airport *domain.Airport) *domain.Airports {
    if _client != nil {
        airports, err := _client.SearchByAirportName(context.Background(), airport)
        if  err != nil {
            log.Errorf("Error: %v\n", err)
        }

        return airports
    }

    return nil
}