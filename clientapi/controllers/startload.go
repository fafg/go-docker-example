package controllers

import (
    "clientapi/application"
    "github.com/labstack/echo/v4"
    log "github.com/sirupsen/logrus"
    "net/http"
    "os"
    "strings"
)

var fileFolderPath = os.Getenv("FILE_FOLDER_PATH")

//StartAirportDatabaseLoad perform the initial load against the grpc server
func StartAirportDatabaseLoad(c echo.Context) error {

    fileName := c.FormValue("filename")
    if strings.TrimSpace(fileName) == "" {
        return c.JSON(http.StatusBadRequest, nil)
    }

    log.Infof("fileName: %s", fileName)
    qty, err := application.LoadData(fileFolderPath + fileName)

    if err != nil {
        return c.JSON(http.StatusNotFound, err)
    }

    return c.JSON(http.StatusCreated, struct { amount int64 }{ qty })
}
