package config

import (
	model "hiringo/model"
)

func Migrate() {
	db := GetDB()

	// Auto Migration
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.Job{})
	db.AutoMigrate(&model.Transaction{})
	db.AutoMigrate(&model.UserDetail{})
	db.AutoMigrate(&model.Rating{})

	CloseDB(db).Close()
}
