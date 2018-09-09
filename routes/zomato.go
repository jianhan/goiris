package routes

import (
	"fmt"
	ghttp "github.com/jianhan/goiris/http"
	"github.com/jianhan/goiris/zomato"
	"github.com/kataras/iris"
	"net/http"
)

func GetZomatoCategoriesHandler(ctx iris.Context) {
	categories, err := zomato.NewCommonAPI().Categories()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("system error, unable initialize client, %s", err.Error()),
				Status:  http.StatusInternalServerError,
			},
		)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(categories)
}
