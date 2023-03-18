package api

import (
	g "QACommunity/global"
	"QACommunity/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/redis/go-redis/v9"
	"log"
)

func InitDB() {
	dns := "root:mysqlsever.1@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	g.Db = db
	g.Db.AutoMigrate(&model.User{}, &model.Question{}, &model.Answer{})

	g.Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
}
