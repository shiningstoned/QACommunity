package global

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/redis/go-redis/v9"
)

var (
	Db       *gorm.DB
	MySecret = []byte("超级加密")
	Client   *redis.Client
	Ctx      = context.Background()
)
