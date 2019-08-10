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

func SortAdd(ctx *gin.Context) {
	// 绑定数据
	data := datamodels.DSort{}
	err := ctx.ShouldBind(&data)
	if err != nil {
		// 返回错误
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	// 入库
	err = repositories.SortAdd(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)
}

func SortModify(ctx *gin.Context) {
	data := datamodels.DSort{}
	err := ctx.ShouldBind(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	err = repositories.SortModify(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)
}

func SortDel(ctx *gin.Context) {
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
	e = repositories.SortDelById(i)
	if e != nil {
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)
}

func SortGetInfoById(ctx *gin.Context) {
	s, b := ctx.Params.Get("id")
	if b != true || s == "" {
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	i, e := strconv.Atoi(s)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	// 删除
	sort, e := repositories.SortSelectById(i)
	if e  != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	data := defs.RequestDate{
		Code:200,
		Data:sort,
	}
	resp.RepMsg(ctx,&data)
}