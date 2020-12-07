package m_client

import (
	"github.com/sumflowns/toyo-src/tree/master/book-ticket-common/basic"
	"github.com/sumflowns/toyo-src/tree/master/book-ticket-common/proto/auth"
	"github.com/micro/go-micro/client"
)

var (
	AuthClient auth.AuthService
)

func Init() {
	AuthClient = auth.NewAuthService(basic.AuthService, client.DefaultClient)
}
