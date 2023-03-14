package util

type EndTime struct {
	endtime string
}

// Monitor 检测mysql数据库，endtime的变化
//func Monitor(){
//	var result []EndTime
//	err := model.DB.Table("charging_pile").Select("end_time").Where("id > ?",0).
//		Scan(&result).Error
//	if err != nil{
//		logging.Info(err)
//	}
//	now := time.Now().UnixMilli()
//	for _ , v := range result{
//		if
//	}
//}
