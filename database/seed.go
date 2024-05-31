package database

import (
	"fmt"
	"log"

	"github.com/2marks/go-expense-tracker-api/types"
	"gorm.io/gorm"
)

var Categories []string = []string{
	"shopping",
	"food",
	"subscription",
	"airtime",
	"groceries",
	"leisure",
	"electronics",
	"utilities",
	"clothing",
	"health",
	"fruit",
	"others",
}

func (d *Database) RunSeeds(db *gorm.DB) {
	if err := seedPaymentCateogories(db); err != nil {
		log.Fatalf("error while running database seeds. err:%s", err.Error())
	}

	log.Println("DB: seeds ran successfully")
}

func seedPaymentCateogories(db *gorm.DB) error {
	for _, category := range Categories {
		var count int64
		db.Model(&types.Category{}).Where("name=? and is_system=1", category).Count(&count)

		if count > 0 {
			continue
		}

		result := db.Create(&types.Category{Name: category, IsSystem: true, CreatedById: 0})
		if result.Error != nil {
			return fmt.Errorf("error while saving category:%s. err:%s", category, result.Error.Error())
		}
	}

	return nil
}
