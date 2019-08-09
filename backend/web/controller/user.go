package controller

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/defs"
	"Vue-Admin/backend/easyutils"
	"Vue-Admin/backend/repositories"
	"Vue-Admin/backend/resp"

	"github.com/gin-gonic/gin"
	"log"
)

// 用户注册
func UserRegister(ctx *gin.Context) {
	// 绑定用户数据
	user := datamodels.DUser{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx, defs.ErrorBadRequest)
		return
	}
	// 入库
	err = repositories.UserRegister(&user)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx, defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx, defs.SuccessOk)
}

// 用户登录
func UserLogin(ctx *gin.Context) {
	// 绑定用户数据
	user := datamodels.DUser{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx, defs.ErrorBadRequest)
		return
	}

	// 验证
	err = repositories.UserCheck(&user)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx, defs.ErrorBadUser)
		return
	}

	// 查询用户是否重复登录 查询
	tok, err := easyutils.EasyJwtCheckUserLogin(user.Email)
	// 如果为空 则 用户已经 登录 返回 当前用户的 token就行了

	// 返回token
	token1 := &defs.RequestDate{
		Code: 200,
		Data: &gin.H{
			"token": tok,
		},
	}

	if err == nil {
		resp.RepMsg(ctx, token1)
		return
	}

	// 生成token
	payload := easyutils.EasyJwtPayload{}
	payload.Iss = user.Email
	s, err := easyutils.EasyJwtGeneraToken(&payload, 8)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx, defs.ErrorBadGeneratingToken)
		return
	}
	// 返回token
	token := &defs.RequestDate{
		Code: 200,
		Data: &gin.H{
			"token": s,
		},
	}
	resp.RepMsg(ctx, token)
}

// 更具 用户email 获取用户数据
func UserInfo(ctx *gin.Context) {
	email, b := ctx.Params.Get("email")
	if !b || email == "" {
		resp.RepMsg(ctx, defs.ErrorBadRequest)
		return
	}
	// 查询用户信息
	user, e := repositories.UserGetInfoByEmail(email)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx, defs.ErrorSQLNotData)
		return
	}
	result := defs.RequestDate{
		Code: 200,
		Data: user,
	}
	resp.RepMsg(ctx, &result)
}
