package business

import (
    "errors"
    log "github.com/sirupsen/logrus"
    "portservice/data"
    "portservice/internal/grpc/domain"
    "strings"
)

func InsertAirport(airport *domain.Airport) error {
    if strings.TrimSpace(airport.Name) == "" {
        log.Error("airport name can't be empty")
        return errors.New("airport name can't be empty")
    }

    return data.InsertAirport(airport)
}

func SearchAirportByName(name string) (*[]*domain.Airport, error) {
    if strings.TrimSpace(name) == "" {
        log.Error("search parameter name can't be empty")
        return nil, errors.New("search parameter name can't be empty")
    }

    return data.SearchAirportByName(name)
}