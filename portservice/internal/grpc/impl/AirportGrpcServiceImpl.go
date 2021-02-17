package impl

import (
    "context"
    "portservice/business"

    log "github.com/sirupsen/logrus"

    "portservice/internal/grpc/domain"
    "portservice/internal/grpc/service"
)

//AirportGrpcServiceImpl is a implementation of AirportService Grpc Service.
type AirportGrpcServiceImpl struct {
}

//NewAirportGrpcServiceImpl returns the pointer to the AirportGrpcServiceImpl implementation.
func NewAirportGrpcServiceImpl() *AirportGrpcServiceImpl {
    return &AirportGrpcServiceImpl{}
}

//Insert function implementation of gRPC Service.
func (serviceImpl *AirportGrpcServiceImpl) Insert(_ context.Context, in *domain.Airport) (*service.InsertAirportResponse, error) {
    log.Infof("adding airport with name: %s\n", in.Codename)

    err := business.InsertAirport(in)
    if err != nil {
        log.Errorf("error on insert airport: %s\n", in.Name)
        return nil, err
    }
    //Logic to persist to data or storage.
    log.Infof("airport %s Added", in.Codename)

    return &service.InsertAirportResponse{
        AddedAirport: in,
        Error:        nil,
    }, nil
}

//Insert function implementation of gRPC Service.
func (serviceImpl *AirportGrpcServiceImpl) BulkInsert(_ context.Context, in *domain.Airports) (*service.BulkInsertAirportResponse, error) {
    log.Infof("adding the quantity of: %v airports\n", len(in.AirportsMap))

    //Logic to persist to data or storage.
    log.Infof("%v airports were added successfully\n", len(in.AirportsMap))

    return &service.BulkInsertAirportResponse{
        AddedAmount: 2,
        Error:        nil,
    }, nil
}

//Insert function implementation of gRPC Service.
func (serviceImpl *AirportGrpcServiceImpl) SearchByAirportName(_ context.Context, in *domain.Airport) (*domain.Airports, error) {
    log.Infof("searching for airport with the name: %s\n", in.Name)

    airports, err := business.SearchAirportByName(in.Name)
    if err != nil {
        log.Errorf("error on insert airport: %s\n", in.Name)
        return nil, err
    }

    //Logic to persist to data or storage.
    log.Infof("%v airports found\n", 3)
    result := make(map[string]*domain.Airport, len(*airports))
    for _, value := range *airports {
        result[value.Codename] = value
    }

    return &domain.Airports{ AirportsMap: result }, nil
}