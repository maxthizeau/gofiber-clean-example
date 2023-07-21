package configuration

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/maxthizeau/gofiber-clean-boilerplate/entity"
	"github.com/maxthizeau/gofiber-clean-boilerplate/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(config Config) *gorm.DB {
	username := config.Get("DATASOURCE_USERNAME")
	password := config.Get("DATASOURCE_PASSWORD")
	host := config.Get("DATASOURCE_HOST")
	port := config.Get("DATASOURCE_PORT")
	dbName := config.Get("DATASOURCE_DB_NAME")
	maxPoolOpen, _ := strconv.Atoi(config.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPoolIdle, _ := strconv.Atoi(config.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPollLifeTime, _ := strconv.Atoi(config.Get("DATASOURCE_POOL_LIFE_TIME"))

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	dsn := "host=" + host + " user=" + username + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicLogging(err)

	psqlDb, err := db.DB()
	exception.PanicLogging(err)

	psqlDb.SetMaxOpenConns(maxPoolOpen)
	psqlDb.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
	psqlDb.SetMaxIdleConns(maxPoolIdle)

	log.Println("Connected to the database successfully")
	log.Println("Running Migrations...")

	err = db.AutoMigrate(&entity.User{}, &entity.Question{}, &entity.Answer{}, &entity.Game{}, &entity.UserRole{})
	exception.PanicLogging(err)

	return db
}
