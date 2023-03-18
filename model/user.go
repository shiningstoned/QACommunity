package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Id       int    `gorm:"primaryKey"`
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type Question struct { //最好自己设置column名称，不然后面查询总要看数据库里的名称
	UserId   int    `gorm:"column:userid"`
	Qid      int    `gorm:"column:qid;auto_increment;primary_key"` //必须设置成主键才能自增？
	QContext string `json:"qcontext" form:"qcontext" gorm:"column:qcontext" binding:"required"`
}

type Answer struct {
	UserId     int    `gorm:"column:userid"`
	QuestionId int    `json:"questionid" form:"questionid" gorm:"column:questionid"`
	AnswerId   int    `json:"answerid" form:"answerid" gorm:"column:answerid;auto_increment;primary_key"`
	AContext   string `json:"acontext" form:"acontext" gorm:"column:acontext" binding:"required"`
}

type MyClaims struct {
	Username string
	jwt.RegisteredClaims
}
