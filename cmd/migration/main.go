package main

import (
	"log"
	"os"
	"strings"

	"github.com/patrickishaf/job_scheduler/config"
	"github.com/patrickishaf/job_scheduler/internal/db"
)

func main() {
	cfg := config.LoadConfig()
	args := os.Args

	if len(args) <= 1 {
		err := db.RunMigrations(cfg.DB.GetConnectionString(), cfg.DB.MigrationsDir)
		if err != nil {
			log.Printf("failed to run migrations::%s", err.Error())
		}
		return
	}
	if strings.ToLower(args[1]) == "down" {
		err := db.RevertMigration(cfg.DB.GetConnectionString(), cfg.DB.MigrationsDir)
		if err != nil {
			log.Printf("failed to run migrations::%s", err.Error())
		}
	} else if strings.ToLower(args[1]) == "reset" {
		err := db.ResetAllMigrations(cfg.DB.GetConnectionString(), cfg.DB.MigrationsDir)
		if err != nil {
			log.Printf("failed to run migrations::%s", err.Error())
		}
	}
}
