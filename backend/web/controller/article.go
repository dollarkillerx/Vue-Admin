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
func ArticleAdd(ctx *gin.Context) {
	data := datamodels.DArticle{}
	err := ctx.ShouldBind(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	err = repositories.ArticleAdd(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx,defs.SuccessOk)
}

// 修改
func ArticleModify(ctx *gin.Context) {
	data := datamodels.DArticle{}
	err := ctx.ShouldBind(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	err = repositories.ArticleModify(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx,defs.SuccessOk)
}

// 删除
func ArticleDel(ctx *gin.Context) {
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

	e = repositories.ArticleDel(i)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)
}

// 更具id 查询
func ArticleGetInfoById(ctx *gin.Context) {
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

	article, e := repositories.ArticleGetById(i)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	data := defs.RequestDate{
		Code: 200,
		Data: article,
	}
	resp.RepMsg(ctx,&data)
}
