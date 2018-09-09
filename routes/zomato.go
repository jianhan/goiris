package routes

import (
	"fmt"
	ghttp "github.com/jianhan/goiris/http"
	"github.com/jianhan/goiris/zomato"
	"github.com/kataras/iris"
	"net/http"
	"os"
)

func GetZomatoCategoriesHandler(ctx iris.Context) {
	commonAPI, err := zomato.NewCommonAPI(os.Getenv("ZOMATO_API_KEY"), os.Getenv("ZOMATO_API_URL"))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("system error, invalid configuration , %s", err.Error()),
				Status:  http.StatusInternalServerError,
			},
		)
		return
	}

	categories, err := commonAPI.Categories()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("system error, unable to retrieve categories, %s", err.Error()),
				Status:  http.StatusInternalServerError,
			},
		)
		return
	}

	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(categories)
}
