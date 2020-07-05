package databases

import (
	"github.com/jinzhu/gorm"
	"github.com/zhuangalbert/boilerplate/src/api/v1/models"
)

func Migrate() {
	db := GetConnection()
	db.AutoMigrate(
		&models.User{},
	)
	ModifyTable(db)
	ModifyIndex(db)
}

func ModifyTable(db *gorm.DB) {

}

func ModifyIndex(db *gorm.DB) {

}
