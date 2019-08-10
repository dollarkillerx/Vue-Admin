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

// 添加
func ArticleItemAdd(ctx *gin.Context) {
	data := datamodels.DArticleItem{}
	err := ctx.ShouldBind(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	err = repositories.ArticleItemAdd(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx,defs.SuccessOk)
}

// 修改
func ArticleItemModify(ctx *gin.Context) {
	data := datamodels.DArticleItem{}
	err := ctx.ShouldBind(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	err = repositories.ArticleItemModify(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx,defs.SuccessOk)
}

// 删除
func ArticleItemDel(ctx *gin.Context) {
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

	e = repositories.ArticleItemDel(i)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)
}

// 通过id 查询
func ArticleItemGetInfoById(ctx *gin.Context) {
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

	item, e := repositories.ArticleItemSelectById(i)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	data := defs.RequestDate{
		Code: 200,
		Data: item,
	}
	resp.RepMsg(ctx,&data)
}
