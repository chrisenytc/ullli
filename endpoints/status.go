package endpoints

import (
	"gopkg.in/kataras/iris.v6"
)

func GetStatus(ctx *iris.Context) {
	ctx.JSON(iris.StatusOK, map[string]string{"message": "Working well."})
}
