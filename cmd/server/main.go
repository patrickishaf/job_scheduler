package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/patrickishaf/job_scheduler/app"
	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/common"
	"github.com/patrickishaf/job_scheduler/internal/db"
	"github.com/patrickishaf/job_scheduler/internal/db/repository"
	"github.com/patrickishaf/job_scheduler/internal/infra"
	"github.com/patrickishaf/job_scheduler/internal/net"
)

func main() {
	cfg := config.LoadConfig()
	logger := common.InitLogger()

	connectionPool, err := db.CreateConnectionPool(context.Background(), cfg.DB.GetConnectionString())
	if err != nil {
		log.Fatalf("failed to create db connection pool. %s", err.Error())
	}

	redisCache := infra.InitRedisCache(&cfg.Redis)
	socketConnStore := net.CreateSocketConnStore()

	jobRepo := repository.CreateJobRepository(&cfg.App, connectionPool, logger)
	dlqRepo := repository.CreateDLQRepository(&cfg.App, connectionPool, logger)

	service := app.CreateService(redisCache, dlqRepo, jobRepo, socketConnStore, logger)

	intervalWorker := app.CreateIntervalWorker(&cfg.App, service, logger)
	scheduledWorker := app.CreateScheduledWorker(&cfg.App, service, logger)
	go intervalWorker.Start()
	go scheduledWorker.Start()

	httpHandler := app.CreateHTTPHandler(service, logger)
	socketHandler := app.CreateSocketHandler(&cfg.Socket, service, logger)
	apiRouter := app.InitRouter(httpHandler, socketHandler)

	srv := app.CreateAppServer(&cfg.App)
	srv.ConfigureRoutes(apiRouter)
	go func() {
		srv.Run()
	}()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	sig := <-exitChan
	logger.Info("signal received. shutting down application", "sig", sig)
	intervalWorker.Stop()
	scheduledWorker.Stop()
	socketConnStore.Clear()
	connectionPool.Close()
}
