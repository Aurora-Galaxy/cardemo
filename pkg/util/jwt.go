package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var JWTsecret = []byte("ABC")

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	jwt.StandardClaims
}

// GenerateToken 签发token
func GenerateToken(id uint, username string, password string) (string, error) {
	NowTime := time.Now()
	expire := NowTime.Add(24 * time.Hour) //token有效期24小时
	claims := Claims{
		Id:       id,
		UserName: username,
		PassWord: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(), //秒级时间戳
			Issuer:    "car"},
	}
	//生成claims对应的token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //SigningMethodHS256 注意密码类型
	token, err := tokenClaims.SignedString(JWTsecret)
	return token, err
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTsecret, nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if ok && tokenClaims.Valid { //tokenClaims.Valid判断token是否被更改过，如果没有就为true
			return claims, nil
		}
	}
	return nil, err
}
