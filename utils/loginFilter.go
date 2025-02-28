package utils

import "github.com/astaxie/beego/context"

func LoginFilter(ctx *context.Context) {
	id := ctx.Input.Query("id")
	if id == "" {
		ctx.Redirect(302, "/")
	}
}
