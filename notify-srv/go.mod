module notify-srv

go 1.13

require (
	github.com/sumflowns/toyo-src/book-ticket-common v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.11
)

replace github.com/sumflowns/toyo-src/book-ticket-common => ../book-ticket-common
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
