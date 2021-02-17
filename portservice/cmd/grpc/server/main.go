package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc/keepalive"
    "net"
    "os"
    "os/signal"
    "portservice/data"
    "syscall"
    "time"

    log "github.com/sirupsen/logrus"
    "golang.org/x/sync/errgroup"
    "google.golang.org/grpc"
    "google.golang.org/grpc/health"
    healthproto "google.golang.org/grpc/health/grpc_health_v1"

    "portservice/internal/grpc/impl"
    "portservice/internal/grpc/service"
    "portservice/logger"
)

var (
    grpcServer       *grpc.Server
    grpcHealthServer *grpc.Server
)

func getServerListener(port uint) net.Listener {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

    if err != nil {
        log.Fatalf("failed to listen: %v", err)
        panic(fmt.Sprintf("failed to listen: %v", err))
    }

    return lis
}

func getHealthServerListener(port uint) net.Listener {
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

    if err != nil {
        log.Fatalf("failed to listen: %v", err)
        panic(fmt.Sprintf("failed to listen: %v", err))
    }

    return lis
}

func main() {
    logger.Init()
    data.Init(os.Getenv("PORT_SERVICE_CONN_STR_DB"))
    data.OpenConnection()
    defer data.CloseConnection()

    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
    defer signal.Stop(signalChan)

    g, ctx := errgroup.WithContext(ctx)

    // gRPC Health Server
    healthServer := health.NewServer()
    g.Go(func() error {
        grpcHealthServer = grpc.NewServer()
        netHealthListener := getHealthServerListener(7071)
        healthproto.RegisterHealthServer(grpcHealthServer, healthServer)

        log.Infof("gRPC health server serving at %s", netHealthListener.Addr().String())

        return grpcHealthServer.Serve(netHealthListener)
    })

    // gRPC Main Server
    g.Go(func() error {
        netListener := getServerListener(7070)
        airportServiceImpl := impl.NewAirportGrpcServiceImpl()

        grpcServer = grpc.NewServer(
            grpc.KeepaliveParams(keepalive.ServerParameters{MaxConnectionAge: 2 * time.Minute}),
        )
        service.RegisterPortDomainServiceServer(grpcServer, airportServiceImpl)

        log.Infof("gRPC server serving at %s", netListener.Addr().String())
        healthServer.SetServingStatus("grpc.health.v1.airportservice", healthproto.HealthCheckResponse_SERVING)

        return grpcServer.Serve(netListener)
    })

    select {
    case <-signalChan:
        break
    case <-ctx.Done():
        break
    }

    log.Warn("Shutdown process has been started.")

    healthServer.SetServingStatus("grpc.health.v1.airportservice", healthproto.HealthCheckResponse_NOT_SERVING)

    if grpcServer != nil {
        log.Info("Shutting down grpc server.")
        grpcServer.GracefulStop()
    }

    if grpcHealthServer != nil {
        log.Info("Shutting down health grpc server.")
        grpcHealthServer.GracefulStop()
    }

    err := g.Wait()
    if err != nil {
        log.Errorf("server returning an error: %v\n", err)
        os.Exit(1)
    }
}
