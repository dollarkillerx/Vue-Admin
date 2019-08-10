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
func CarouselAdd(ctx *gin.Context) {
	data := datamodels.DCarousel{}
	err := ctx.ShouldBind(&data)

	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	err = repositories.CarouselAdd(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx,defs.SuccessOk)
}

// 修改
func CarouselModify(ctx *gin.Context) {
	data := datamodels.DCarousel{}
	err := ctx.ShouldBind(&data)

	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	err = repositories.CarouselModify(&data)
	if err != nil {
		log.Println(err.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}
	resp.RepMsg(ctx,defs.SuccessOk)
}

// 删除
func CartouselDelById(ctx *gin.Context) {
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

	e = repositories.CarouselDel(i)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	resp.RepMsg(ctx,defs.SuccessOk)

}

// 通过id 查询
func CarouselGetInfoById(ctx *gin.Context) {
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
	carousel, e := repositories.CarouseSelectById(i)
	if e != nil {
		log.Println(e.Error())
		resp.RepMsg(ctx,defs.ErrorBadRequest)
		return
	}

	data := defs.RequestDate{
		Code: 200,
		Data: carousel,
	}
	resp.RepMsg(ctx,&data)
}