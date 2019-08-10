package controller

import (
	"Vue-Admin/backend/datamodels"
	"Vue-Admin/backend/defs"
	"Vue-Admin/backend/repositories"
	"Vue-Admin/backend/resp"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// 添加导航
func NavAdd(ctx *gin.Context) {
	nav := datamodels.DNav{}

	err := ctx.ShouldBind(&nav)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	// 入库
	err = repositories.NavAdd(&nav)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)

}

// 修改导航
func NavModify(ctx *gin.Context) {
	nav := datamodels.DNav{}
	err := ctx.ShouldBind(&nav)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	// modify
	err = repositories.NavModify(&nav)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)
}

// 删除导航
func NavDel(ctx *gin.Context) {
	s, b := ctx.Params.Get("id")
	if b != true || s == "" {
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	i, e := strconv.Atoi(s)
	if e != nil {
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	// 删除
	err := repositories.NavDel(i)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx,defs.SuccessOk)
}

// 通过id 查询
func NavGetInfoById(ctx *gin.Context) {
	s, b := ctx.Params.Get("id")
	if b != true || s == "" {
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	i, e := strconv.Atoi(s)
	if e != nil {
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	// 查询
	nav, err := repositories.NavSelectById(i)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	data := defs.RequestDate{
		Code: 200,
		Data: nav,
	}
	resp.RepMsg(ctx,&data)
}