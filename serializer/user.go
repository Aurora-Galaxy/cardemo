package serializer

import "car/model"

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CarNumber string `json:"car_number"`
	//CreateAt int64  `json:"create_at"`
}

//BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		UserName:  user.UserName,
		Phone:     user.Phone,
		Email:     user.Email,
		CarNumber: user.CarNUmber,
		//CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
