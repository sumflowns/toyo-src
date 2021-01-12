package jwt

import "github.com/sumflowns/toyo-src/book-ticket-common/basic"

// Jwt 配置 接口
type Jwt struct {
	SecretKey string `json:"secretKey"`
}

// init 初始化Redis
func init() {
	basic.Register(initJwt)
}

func initJwt() {

}
