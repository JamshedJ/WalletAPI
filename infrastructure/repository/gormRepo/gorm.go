package gormRepo

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(dsn string) error {
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	return err
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&gormWallet{},
	)
}

func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
