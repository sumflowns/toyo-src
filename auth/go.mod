module surprise-shop-auth

go 1.13

require (
github.com/sumflowns/toyo-src/tree/master/book-ticket-common v0.0.0-00010101000000-000000000000
github.com/dgrijalva/jwt-go v3.2.0+incompatible
github.com/go-redis/redis v6.15.6+incompatible
github.com/micro/cli v0.2.0
github.com/micro/go-micro v1.17.1
github.com/micro/go-plugins v1.5.1
github.com/opentracing/opentracing-go v1.1.0
)

replace github.com/sumflowns/toyo-src/tree/master/book-ticket-common => ../book-ticket-common
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
