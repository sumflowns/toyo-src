package main

import (
	_ "github.com/sumflowns/toyo-src/tree/master/book-ticket-common/plugins/db"
	//	_ "github.com/sumflowns/toyo-src/tree/master/book-ticket-common/plugins/redis"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/micro/go-plugins/broker/nats"
)
