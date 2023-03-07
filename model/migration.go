package model

//数据自动迁移

func Migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).
		AutoMigrate(&ChargingPile{})
	//.AutoMigrate(&Car{}).AutoMigrate(&ChargingPile{})
}
