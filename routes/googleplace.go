package routes

import "github.com/kataras/iris"

func GetGooglePlaceHandler(ctx iris.Context) {
	ctx.Write([]byte("test"))
}
