package controller

import (
	"github.com/mzz2017/v2rayA/common"
	"github.com/mzz2017/v2rayA/db/configure"
	"github.com/mzz2017/v2rayA/service"
	"github.com/gin-gonic/gin"
)

func PostImport(ctx *gin.Context) {
	var data struct {
		URL   string `json:"url"`
		Which *configure.Which
	}
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		common.ResponseError(ctx, logError(nil, "bad request"))
		return
	}
	err = service.Import(data.URL, data.Which)
	if err != nil {
		common.ResponseError(ctx, logError(err))
		return
	}
	GetTouch(ctx)
}
