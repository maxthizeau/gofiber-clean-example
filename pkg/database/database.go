package database

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/maxthizeau/gofiber-clean-boilerplate/pkg/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConfig struct {
	User            string
	Password        string
	Host            string
	Port            string
	DBName          string
	MaxPoolOpen     int
	MaxPoolIdle     int
	MaxPollLifeTime int
}

func NewDatabase(config DatabaseConfig) *gorm.DB {

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	dsn := "host=" + config.Host + " user=" + config.User + " password=" + config.Password + " dbname=" + config.DBName + " port=" + config.Port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicLogging(err)

	psqlDb, err := db.DB()
	exception.PanicLogging(err)

	psqlDb.SetMaxOpenConns(config.MaxPoolOpen)
	psqlDb.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(config.MaxPollLifeTime))) * time.Millisecond)
	psqlDb.SetMaxIdleConns(config.MaxPoolIdle)

	log.Println("Connected to the database successfully")

	// log.Println("Running Migrations...")
	// err = db.AutoMigrate(&entity.User{}, &entity.Question{}, &entity.Answer{}, &entity.Game{}, &entity.UserRole{})
	// exception.PanicLogging(err)

	return db
}
