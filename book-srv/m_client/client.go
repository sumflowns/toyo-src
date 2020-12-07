package m_client

import (
	"github.com/sumflowns/toyo-src/tree/master/book-ticket-common/basic"
	user "github.com/sumflowns/toyo-src/tree/master/book-ticket-common/proto/user"
	"github.com/micro/go-micro/client"
)

func Init() {
	UserClient = user.NewUserInfoService(basic.UserService, client.DefaultClient)
}

var (
	UserClient user.UserInfoService
)
