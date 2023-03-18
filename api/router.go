package api

import (
	"QACommunity/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	r := gin.Default()
	r.POST("/register", Register) //注册
	r.POST("/login", Login)       //登录
	r1 := r.Group("/user")
	{
		r1.Use(middleware.AuthToken())              //验证token
		r1.POST("/inquiry", Inquiry)                //发布问题
		r1.POST("/answer", Answer)                  //回答问题
		r1.GET("/getquestions", GetQuestions)       //获取问题
		r1.GET("/getanswers", GetAnswers)           //获取回答
		r1.POST("/getanswers/like", LikeAnswer)     //给问题点赞
		r1.POST("/getanswers/unlike", UnlikeAnswer) //取消点赞
		r1.POST("/following", Following)            //关注用户
		r1.GET("showfollowing", ShowFollowing)      //查看关注
		r1.GET("showfollowers", ShowFollowers)      //查看粉丝
		r1.DELETE("/deleteanswer", DeleteAnswer)    //删除回答
		r1.DELETE("deletequestion", DeleteQuestion) //删除问题
		r1.PUT("/modifyanswer", ModifyAnswer)       //修改回答
		r1.PUT("/modifyquestion", ModifyQusetion)   //修改问题
	}
	r.Run(":8888")

}
