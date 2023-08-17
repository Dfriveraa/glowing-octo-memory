package repositories

import (
	"github.com/dfriveraa/glowing-octo-memory/app/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Db struct {
	client *gorm.DB
}

func InitDB() *Db {
	var err error
	dsn := "user=memory password=K4T7VpOXrDxe dbname=octomemory host=ep-royal-star-125567.us-west-2.aws.neon.tech sslmode=verify-full"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&domain.User{})

	if err != nil {
		panic(err)
	}

	if err != nil {
		panic("failed to migrate database")
	}
	return &Db{client: database}
}

func (db *Db) Close() error {
	sqlDB, err := db.client.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
