package defs

import "github.com/gin-gonic/gin"

type RequestDate struct {
	Code int
	Data interface{}
}

// 错误
var (
	ErrorBadRequest         = &RequestDate{Code: 400, Data: &gin.H{"Code": "40001", "Msg": "Bad required"}}
	ErrorSQLNotData         = &RequestDate{Code: 404, Data: &gin.H{"Code": "40401", "Msg": "SQL Not Found"}}
	ErrorBadServer          = &RequestDate{Code: 500, Data: &gin.H{"Code": "50001", "Msg": "Internal Server Error"}}
	ErrorBadUser            = &RequestDate{Code: 400, Data: &gin.H{"Code": "40002", "Msg": "Error in username or password"}}
	ErrorBadGeneratingToken = &RequestDate{Code: 400, Data: &gin.H{"Code": "40003", "Msg": "Failed to generate token"}}
	ErrorUnauthorized       = &RequestDate{Code: 401, Data: &gin.H{"Code": "40010", "Msg": "401 Unauthorized"}}
)

// 正确
var (
	SuccessOk = &RequestDate{Code: 200, Data: &gin.H{"Code": "20001", "Msg": "Success Ok"}}
)
