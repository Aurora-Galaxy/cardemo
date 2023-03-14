package model

//数据自动迁移

func Migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&User{}).
		AutoMigrate(&ChargingPile{}).AutoMigrate(&HistoryRecord{})
	//.AutoMigrate(&Car{}).AutoMigrate(&ChargingPile{})
	DB.Model(&HistoryRecord{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE") //CASCADE跟随外键改动
}
