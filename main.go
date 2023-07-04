package main

import (
	"fmt"
	"os"
	"time"

	_ "github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/docs"
	"github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/internal/configs"
	"github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/internal/db/repositories"
	grpcserver "github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/internal/grpc_server"
	"github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/internal/services"

	"github.com/rs/zerolog/log"

	"github.com/go-co-op/gocron"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	env := os.Getenv("environment")
	if env == "" {
		env = "dev"
	}

	//Load configuration
	conf, confErr := configs.LoadConfigs(env)
	if confErr != nil {
		log.Fatal().Err(confErr).Msg("Config load failed")
	}

	//DB Init
	dbConfig := conf.PostgresDB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed")
	}

	//Setuping sync job
	rep := repositories.NewCurrencySnapshotDataService(db)
	syncService := services.NewSyncService(rep)
	job := gocron.NewScheduler(time.UTC)
	job.Every(1).Day().At(conf.Server.SyncTime1).At(conf.Server.SyncTime2).Do(func() {
		syncService.StartSync()
	})

	job.StartAsync()

	//// Server REST Init
	//server.Init(db)

	// Server gRPC Init
	grpcserver.Init(db)
}
