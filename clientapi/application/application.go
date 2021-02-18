package application

import (
    "encoding/json"
    "errors"
    "os"
    "strings"

    log "github.com/sirupsen/logrus"

    "clientapi/grpc"
    "clientapi/internal/grpc/domain"
)

func LoadData(jsonFile string) (int64, error) {
    if fileExist(jsonFile) {
        var amount int64
        file, _ := os.Open(jsonFile)
        defer func () {
            err := file.Close()
            if err != nil {
                log.Errorf("error on trying to close the json file: %s\n", err)
            }
        }()

        decoder := json.NewDecoder(file)
        var airport = domain.Airport{}

        for decoder.More() {
            token, _ := decoder.Token()
            log.Infof("token: %v\n", token)

            err := decoder.Decode(&airport)
            if err == nil {
                airport.Codename = token.(string)
                response := grpc.SendAirportToRegistration(&airport)
                if response != nil && response.Error == nil {
                    amount = amount + 1
                }
            } else {
                log.Warnf("error on decoding json item: %s\n", err)
            }
        }

        return amount, nil
    }

    return 0, errors.New("file not found")
}

func GetAirportByName(airport *domain.Airport) *domain.Airports {
    if strings.TrimSpace(airport.Name) == "" {
        return nil
    }

    return grpc.SearchAirportByName(airport)
}

func fileExist(filename string) bool {
    if _, err := os.Stat(filename); err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }

    return true
}