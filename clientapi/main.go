package main

import (
    "clientapi/grpc"
    "clientapi/logger"
    "clientapi/server"
    "context"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    logger.Init()

    newServer := server.NewServer()
    go func() {
        newServer.Start(":8000")
    }()

    grpc.Init()
    defer grpc.Close()

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
    defer signal.Stop(signalChan)

    select {
    case <-signalChan:
        break
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
    defer cancel()
    newServer.Shutdown(ctx)
}
