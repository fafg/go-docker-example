package data

import (
    "portservice/internal/grpc/domain"
    //"github.com/afex/hystrix-go/hystrix"
)

type AirportDatabase interface {
    Init(connStr string)
    OpenConnection()
    CloseConnection()

    InsertAirport(airport *domain.Airport) error
    SearchAirportByName(name string) (*[]*domain.Airport, error)
}

var dbImpl *MongoDb

func Init(connStr string) {
    dbImpl = newMongoDb()
    dbImpl.Init(connStr)
}

func OpenConnection() {
    dbImpl.OpenConnection()
}

func CloseConnection() {
    dbImpl.CloseConnection()
}

func InsertAirport(airport *domain.Airport) error {
    return dbImpl.InsertAirport(airport)
}

func SearchAirportByName(name string) (*[]*domain.Airport, error) {
    return dbImpl.SearchAirportByName(name)
}