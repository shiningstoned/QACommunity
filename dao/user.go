package dao

import (
	g "QACommunity/global"
	"QACommunity/model"
)

func MatchId(username any) int {
	var user model.User
	g.Db.Find(&user, "username=?", username)
	//fmt.Println(user.ID)
	return user.Id
}

func MatchName(id any) string {
	var user model.User
	g.Db.Find(&user, "id=?", id)
	return user.Username
}
