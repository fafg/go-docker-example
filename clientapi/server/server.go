package server

import (
    "context"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    log "github.com/sirupsen/logrus"

    "clientapi/controllers"
)

// StockAPI structur
type StockAPI struct {
    echoApp *echo.Echo
}

//NewServer Instance of Echo
func NewServer() *StockAPI {
    return &StockAPI{
        echoApp: echo.New(),
    }
}

//Start server functionality
func (server *StockAPI) Start(port string) {
    // logger
    server.echoApp.Use(middleware.Logger())
    // recover
    server.echoApp.Use(middleware.Recover())
    //CORS
    server.echoApp.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{echo.GET, echo.POST},
    }))

    // load endpoint
    server.echoApp.POST("/v1/clientapi/startload", controllers.StartAirportDatabaseLoad)
    // search endpoint
    server.echoApp.GET("/v1/clientapi/search/:name", controllers.GetAirportByName)

    // Start Server
    err := server.echoApp.Start(port)
    if err != nil {
        log.Warnf("%v\n", err)
    }
}

//Shutdown server functionality
func (server *StockAPI) Shutdown(ctx context.Context) {
    err := server.echoApp.Shutdown(ctx)
    if err != nil {
        log.Errorf("error on try to close echo application server: %v\n", err)
    }
}

