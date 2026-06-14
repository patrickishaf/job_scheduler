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

	heapScheduler := app.NewHeapScheduler()
	ipqScheduler := app.NewIndexedPQScheduler()

	worker := app.CreateWorker(&cfg.App, heapScheduler, ipqScheduler, jobRepo, socketConnStore, logger)
	queueWorker := app.CreateQueueWorker(&cfg.App, heapScheduler, ipqScheduler, socketConnStore, logger)
	workerCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go worker.Start(workerCtx)
	go queueWorker.Start(workerCtx)

	httpHandler := app.CreateHTTPHandler(service, logger)
	socketHandler := app.CreateSocketHandler(&cfg.Socket, service, logger)
	apiRouter := app.InitRouter(httpHandler, socketHandler)

	srv := app.CreateAppServer(&cfg.App)
	srv.ConfigureRoutes(apiRouter)
	go srv.Run()

	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	sig := <-exitChan
	logger.Info(common.SHUTTING_DOWN_SERVER, "signal", sig)
	worker.Stop()
	queueWorker.Stop()
	socketConnStore.Clear()
	connectionPool.Close()
}
