package controllers

import (
    "clientapi/application"
    "clientapi/internal/grpc/domain"
    "github.com/labstack/echo/v4"
    log "github.com/sirupsen/logrus"
    "net/http"
    "net/url"
    "strings"
)

//GetAirportByName perform a search by name against grpc server
func GetAirportByName(c echo.Context) (err error) {
    // airport name from path `search/:name`
    name, _ := url.QueryUnescape(c.Param("name"))

    if strings.TrimSpace(name) == "" {
        return c.JSON(http.StatusBadRequest, nil)
    }

    log.Infof("name: %s", name)
    airports := application.GetAirportByName(&domain.Airport{Name: name})

    return c.JSON(http.StatusOK, airports.AirportsMap) //also could be 302 - Found
}
