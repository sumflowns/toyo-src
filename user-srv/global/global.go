package global

import (
	auth "github.com/sumflowns/toyo-src/book-ticket-common/proto/auth"
	r "github.com/go-redis/redis"
)

var (
	RedisClient *r.Client
	AuthClient  auth.AuthService
)
