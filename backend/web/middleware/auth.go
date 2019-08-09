package middleware

import (
	"Vue-Admin/backend/defs"
	"Vue-Admin/backend/easyutils"
	"Vue-Admin/backend/resp"
	"github.com/gin-gonic/gin"
)

// 权鉴
func Auth(ctx *gin.Context) {
	// 获取token
	token := ctx.Request.Header.Get("token")
	// 判断token 是否存在
	if token == "" {
		// 如果用户没有携带token 这返回401
		resp.RepMsg(ctx, defs.ErrorBadGeneratingToken)
		ctx.Abort()
		return
	}

	// 如果用户携带token 验证token是否正确 如果正确这通过 反之返回401

	// 验证token
	verification := easyutils.EasyJwtVerification(token)
	if verification != nil {
		//  这返回401
		resp.RepMsg(ctx, defs.ErrorBadGeneratingToken)
		ctx.Abort()
		return
	}
	// 没有错误 通过验证


}
