package middleware

import (
	"QACommunity/service"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthToken() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(200, gin.H{
				"msg": "没有token",
			})
			c.Abort()
			return
		}
		arr := strings.Split(token, " ")
		//fmt.Println(arr)
		tokenString := arr[1]
		//fmt.Println(tokenString)
		claims, err := service.ParseToken(tokenString)
		//fmt.Println(claims)
		if err != nil {
			c.JSON(200, gin.H{"msy": "解析token失败"})
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
