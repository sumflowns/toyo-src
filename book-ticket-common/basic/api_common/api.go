package api_common

import (
	"fmt"
	"github.com/sumflowns/toyo-src/book-ticket-common/basic/utils/string_utl"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/util/log"
	"net/http"
)

type ResponseEntity struct {
	Code    int64       `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ListResponseEntity struct {
	Code    int64       `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Total   int64       `json:"total,omitempty"`
	Page    int64       `json:"page,omitempty"`
	Limit   int64       `json:"limit,omitempty"`
}

type Error struct {
	Code    int32
	Message string
}

func SrvResultListDone(c *gin.Context, data interface{}, limit int64, pages int64, total int64, srvErr *Error) {
	response := &ListResponseEntity{}
	if srvErr.Code == http.StatusInternalServerError {
		log.Logf("Err: %v", srvErr)
		response.Message = "服务端异常，请稍后再试"
		response.Code = http.StatusInternalServerError
		c.AbortWithStatusJSON(http.StatusOK, response)
	}
	if srvErr.Code != http.StatusOK {
		response.Message = srvErr.Message
		response.Code, _ = string_utl.StringToInt64(fmt.Sprintf("%d", srvErr.Code))
		c.AbortWithStatusJSON(http.StatusOK, response)
	} else {
		response.Message = srvErr.Message
		response.Code = http.StatusOK
		response.Data = data
		response.Page = pages
		response.Total = total
		response.Limit = limit
		c.AbortWithStatusJSON(http.StatusOK, response)
	}
}

func SrvResultDone(c *gin.Context, data interface{}, srvErr *Error) {
	response := &ResponseEntity{}
	if srvErr.Code == http.StatusInternalServerError {
		log.Logf("Err: %v", srvErr)
		response.Message = "服务端异常，请稍后再试"
		response.Code = http.StatusInternalServerError
		c.AbortWithStatusJSON(http.StatusOK, response)
	}
	if srvErr.Code != http.StatusOK {
		response.Message = srvErr.Message
		response.Code, _ = string_utl.StringToInt64(fmt.Sprintf("%d", srvErr.Code))
		c.AbortWithStatusJSON(http.StatusOK, response)
	} else {
		response.Message = srvErr.Message
		response.Code = http.StatusOK
		response.Data = data
		c.AbortWithStatusJSON(http.StatusOK, response)
	}
}

//获取用户ID
func GetHeadUserId(c *gin.Context) (userId string) {
	userIdString := c.Request.Header.Get("userId")
	return userIdString
}