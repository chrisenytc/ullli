package endpoints

import (
	"gopkg.in/kataras/iris.v6"
)

func GetIndex(ctx *iris.Context) {
	ctx.MustRender("index.html", nil)
}
