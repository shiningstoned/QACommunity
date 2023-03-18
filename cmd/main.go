package main

import (
	"QACommunity/api"
	g "QACommunity/global"
)

func main() {
	api.InitDB()
	defer g.Db.Close()
	api.SetupRouter()
}
