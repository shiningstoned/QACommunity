package api

import (
	"QACommunity/dao"
	g "QACommunity/global"
	"QACommunity/model"
	"QACommunity/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "格式有误",
		})
		return
	}
	flag := service.IsNameExist(user.Username)
	if flag {
		c.JSON(200, gin.H{"msg": "username already exist"})
		return
	}
	if err := g.Db.Create(&user).Error; err != nil {
		c.JSON(200, gin.H{"msg": "failed register"})
	} else {
		c.JSON(200, gin.H{"msg": "register successfully"})
	}
}

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(200, gin.H{"err": err.Error()})
		return
	}
	flag := service.IsNameExist(user.Username)
	if !flag {
		c.String(200, "user dose not exist")
		return
	} else {
		c.String(200, "login successfully\n")
	}
	TokenString, err := service.SetToken(user.Username)
	if err != nil {
		c.String(200, "I don't know what to say")
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "here is the token",
		"token": TokenString,
	})
}

func Inquiry(c *gin.Context) {
	var question model.Question
	question.QContext = c.PostForm("qcontext")
	if question.QContext == "" {
		c.JSON(200, gin.H{"msg": "问题不能为空"})
		return
	}
	username, ok := c.Get("username")
	if !ok {
		c.JSON(200, gin.H{"msg": "登录已过期"})
	}
	Uid := dao.MatchId(username)
	question.UserId = Uid
	result := g.Db.Create(&question)
	if result.RowsAffected != 0 {
		c.JSON(200, gin.H{
			"msg": "问题发布成功",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "问题发布失败",
		})
	}
}

func Answer(c *gin.Context) {
	var answer model.Answer
	username, _ := c.Get("username")
	Uid := dao.MatchId(username)
	answer.UserId = Uid
	questionid := c.PostForm("questionid")
	Acontext := c.PostForm("acontext")
	if Acontext == "" {
		c.JSON(200, gin.H{"msg": "回答不能为空"})
		return
	}
	answer.AContext = Acontext
	answer.QuestionId, _ = strconv.Atoi(questionid) //将string转为int
	result := g.Db.Create(&answer)
	if result.RowsAffected != 0 {
		c.JSON(200, gin.H{"msg": "回答问题成功"})
	} else {
		c.JSON(200, gin.H{"msg": "回答问题失败"})
	}
}

func GetQuestions(c *gin.Context) {
	var questions []model.Question
	username, ok := c.Get("username")
	if !ok {
		c.JSON(200, gin.H{"msg": "token已失效"})
		return
	}
	Uid := dao.MatchId(username)
	g.Db.Where("userid=?", Uid).Find(&questions)
	for _, question := range questions {
		c.JSON(200, gin.H{"msg": question.QContext})
	}
	//fmt.Println(questions)
	//c.JSON(200, string("hello"))
}

func GetAnswers(c *gin.Context) {
	questionid := c.PostForm("questionid")
	var answers []model.Answer
	result := g.Db.Find(&answers, "questionid=?", questionid)
	if result.RowsAffected != 0 {
		for _, answer := range answers {
			c.JSON(200, gin.H{"msg": answer.AContext})
		}
	} else {
		c.JSON(200, gin.H{"msg": "获取问题失败"})
	}
}

func DeleteAnswer(c *gin.Context) {
	answerid := c.PostForm("answerid")
	var answer model.Answer
	g.Db.Where("answerid=?", answerid).Find(&answer)
	username, _ := c.Get("username")
	Uid := dao.MatchId(username)
	if answer.UserId != Uid {
		c.JSON(200, gin.H{"msg": "这不是你的回答"})
		return
	} else {
		g.Db.Delete(&answer)
		c.JSON(200, gin.H{"msg": "成功删除回答"})
	}
}

func DeleteQuestion(c *gin.Context) {
	questionid := c.PostForm("questionid")
	var question model.Question
	g.Db.Where("qid=?", questionid).Find(&question)
	username, _ := c.Get("username")
	Uid := dao.MatchId(username)
	if question.UserId != Uid {
		c.JSON(200, gin.H{"msg": "这不是你的问题"})
		return
	} else {
		flag := service.DeleteAllAnswers(questionid, Uid)
		if !flag {
			c.JSON(200, gin.H{"msg": "删除回答失败"})
		} else {
			g.Db.Delete(&question)
			c.JSON(200, gin.H{"msg": "问题及回答成功删除"})
		}
	}
}

func ModifyAnswer(c *gin.Context) {
	username, _ := c.Get("username")
	Uid := dao.MatchId(username)
	answerid := c.PostForm("answerid")
	//fmt.Println(username, Uid, answerid)
	var answer model.Answer
	g.Db.Where("answerid=?", answerid).Find(&answer)
	//fmt.Println(Uid, answer)
	if answer.UserId != Uid {
		c.JSON(200, gin.H{"msg": "这不是你的回答"})
		return
	} else {
		acontext := c.PostForm("acontext")
		fmt.Println(acontext)
		result := g.Db.Model(&model.Answer{}).Where("answerid=?", answerid).Update("acontext", acontext)
		if result.RowsAffected == 0 {
			c.JSON(200, gin.H{"msg": "修改问题失败"})
			return
		} else {
			c.JSON(200, gin.H{"msg": "修改问题成功"})
		}
	}
}

func ModifyQusetion(c *gin.Context) {
	username, _ := c.Get("username")
	Uid := dao.MatchId(username)
	questionid := c.PostForm("questionid")
	var question model.Question
	g.Db.Where("qid=?", questionid).Find(&question)
	if question.UserId != Uid {
		c.JSON(200, gin.H{"msg": "这不是你的问题"})
		return
	} else {
		qcontext := c.PostForm("qcontext")
		result := g.Db.Model(&model.Question{}).Where("qid=?", questionid).Update("qcontext", qcontext)
		if result.RowsAffected == 0 {
			c.JSON(200, gin.H{"msg": "修改问题失败"})
		} else {
			c.JSON(200, gin.H{"msg": "修改问题成功"})
		}
	}
}

func LikeAnswer(c *gin.Context) {
	answerId := c.PostForm("answerid")
	username, _ := c.Get("username")
	userId := dao.MatchId(username)
	err := g.Client.SAdd(context.Background(), answerId, userId).Err()
	if err != nil {
		c.JSON(200, gin.H{"err": err.Error()})
		return
	}
	count, _ := g.Client.SCard(g.Ctx, answerId).Result()
	c.JSON(200, gin.H{
		"msg":   "点赞成功",
		"likes": count,
	})
}

func UnlikeAnswer(c *gin.Context) {
	answerId := c.PostForm("answerid")
	username, _ := c.Get("username")
	userId := dao.MatchId(username)
	err := g.Client.SRem(g.Ctx, answerId, userId).Err()
	if err != nil {
		c.JSON(200, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "已取消点赞"})
}

func Following(c *gin.Context) {
	beFollowedName := c.PostForm("username")
	beFollowId := dao.MatchId(beFollowedName)
	befollowid := strconv.Itoa(beFollowId)
	myName, _ := c.Get("username")
	myId := dao.MatchId(myName)
	myid := strconv.Itoa(myId)
	g.Client.SAdd(g.Ctx, myid+":following", befollowid)
	g.Client.SAdd(g.Ctx, befollowid+":followers", myid)
	c.JSON(200, gin.H{"msg": "关注成功"})
}

func ShowFollowing(c *gin.Context) {
	myName, _ := c.Get("username")
	myId := dao.MatchId(myName)
	myid := strconv.Itoa(myId)
	count, _ := g.Client.SCard(g.Ctx, myid+":following").Result()
	if count == 0 {
		c.JSON(200, gin.H{"msg": "你暂未关注用户"})
		return
	}
	following, err := g.Client.SMembers(g.Ctx, myid+":following").Result()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"msg": "你关注的用户如下"})
	for _, fl := range following {
		name := dao.MatchName(fl)
		c.JSON(200, gin.H{"msg": name})
	}
}

func ShowFollowers(c *gin.Context) {
	myName, _ := c.Get("username")
	myId := dao.MatchId(myName)
	myid := strconv.Itoa(myId)
	count, _ := g.Client.SCard(g.Ctx, myid+":followers").Result()
	if count == 0 {
		c.JSON(200, gin.H{"msg": "你还没有粉丝"})
		return
	}
	followers, err := g.Client.SMembers(g.Ctx, myid+":followers").Result()
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"msg": "你的粉丝如下"})
	for _, flers := range followers {
		name := dao.MatchName(flers)
		c.JSON(200, gin.H{"msg": name})
	}
}
