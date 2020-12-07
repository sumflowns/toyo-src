package model

import (
	"book-user_srv/global"
	"book-user_srv/model/user_info"
	"github.com/sumflowns/toyo-src/tree/master/book-ticket-common/basic"
	"github.com/sumflowns/toyo-src/tree/master/book-ticket-common/plugins/redis"
	"github.com/sumflowns/toyo-src/tree/master/book-ticket-common/proto/auth"
	"github.com/micro/go-micro/client"
)

func Init() {
	user_info.Init()
	global.RedisClient = redis.Redis()
	global.AuthClient = auth.NewAuthService(basic.AuthService, client.DefaultClient)

}
