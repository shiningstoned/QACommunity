package service

import (
	g "QACommunity/global"
	"QACommunity/model"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func IsNameExist(username string) bool {
	var user model.User
	g.Db.Find(&user, "username = ?", username)
	if user.Username != "" {
		return true //用户名已存在
	}
	return false
}

func SetToken(username string) (tokenString string, err error) {
	claim := model.MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "local.org",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err = token.SignedString(g.MySecret)
	return tokenString, err
}

func ParseToken(tokenString string) (*model.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, Secret())
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("can't handle this token")
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return g.MySecret, nil
	}
}

func DeleteAllAnswers(questionid string, Uid int) bool {
	result := g.Db.Where("questionid=? AND userid=?", questionid, Uid).Delete(&model.Answer{}) //不加条件全部删除？
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
