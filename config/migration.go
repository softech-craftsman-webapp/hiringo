package config

func Migrate() {
	db := GetDB()

	// Auto Migration
	// db.AutoMigrate(&model.File{})

	CloseDB(db).Close()
}
