package logger

import (
    "os"

    "github.com/keepeye/logrus-filename"
    log "github.com/sirupsen/logrus"
)

func Init() {
    filenameHook := filename.NewHook()
    filenameHook.Field = "line"
    log.AddHook(filenameHook)

    log.SetFormatter(&log.TextFormatter{
        FullTimestamp: true,
        DisableSorting: true,
    })

    log.SetOutput(os.Stdout) //this avoid stderr
    log.SetLevel(log.InfoLevel)
}