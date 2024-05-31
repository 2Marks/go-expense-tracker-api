package database

import (
	"fmt"
	"log"

	"github.com/2marks/go-expense-tracker-api/config"
	"github.com/2marks/go-expense-tracker-api/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Envs.DbUser,
		config.Envs.DbPassword,
		config.Envs.DbAddress,
		config.Envs.DbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("could not connect to database: %s", err.Error())
	}

	log.Println("DB: connected successfully")

	return db
}

func (d *Database) RunMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&types.User{},
		&types.PasswordResetLog{},
		&types.Expense{},
		&types.Category{},
	)

	if err != nil {
		log.Fatalf("error while running database migrations: %s", err.Error())
	}

	log.Println("DB: migration ran successfully")
}
