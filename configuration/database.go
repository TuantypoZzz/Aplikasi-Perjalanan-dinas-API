package configuration

import (
	"fmt"
	"golang-todo-app/entity"
	"golang-todo-app/exception"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {
	username := os.Getenv("DATASOURCE_USERNAME")
	password := os.Getenv("DATASOURCE_PASSWORD")
	host := os.Getenv("DATASOURCE_HOST")
	port := os.Getenv("DATASOURCE_PORT")
	dbName := os.Getenv("DATASOURCE_DB_NAME")

	maxPoolOpen, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_MAX_CONN"))
	exception.PanicLogging(err)
	maxPoolIdle, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_IDLE_CONN"))
	exception.PanicLogging(err)
	maxPollLifeTime, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_LIFE_TIME"))
	exception.PanicLogging(err)

	loggerDb := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: loggerDb,
	})
	exception.PanicLogging(err)

	sqlDB, err := db.DB()
	exception.PanicLogging(err)

	sqlDB.SetMaxOpenConns(maxPoolOpen)                                                              //pengaturan berapa jumlah koneksi maksimal yang dibuat
	sqlDB.SetMaxIdleConns(maxPoolIdle)                                                              //pengaturan berapa jumlah koneksi minimal yang dibuat
	sqlDB.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond) //pengaturan berapa lama koneksi boleh digunakan

	//autoMigrate
	err = db.AutoMigrate(
		&entity.User{},
		&entity.Role{},
		&entity.Permission{},
		&entity.Todo{},
		&entity.Department{},
		&entity.Handle{},
		&entity.Employee{},
		&entity.BussinessTravelReport{},
		&entity.PerdinEmployee{},
		&entity.Lumpsum{},
		&entity.Transport{},
		&entity.Accommodation{},
		&entity.OtherCostDetail{},
		&entity.Telaah{},
		&entity.Spt{},
		&entity.Sppd{},
		&entity.Activity{},
		&entity.DokProofPerdin{},
	)
	exception.PanicLogging(err)

	err = db.SetupJoinTable(&entity.Employee{}, "BussinessTravelReports", &entity.PerdinEmployee{})
	exception.PanicLogging(err)

	err = db.SetupJoinTable(&entity.BussinessTravelReport{}, "Employees", &entity.PerdinEmployee{})
	exception.PanicLogging(err)

	return db
}
