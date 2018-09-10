package routes

import (
	"fmt"
	ghttp "github.com/jianhan/goiris/http"
	"github.com/jianhan/goiris/zomato"
	"github.com/kataras/iris"
	"net/http"
)

type zomatoRoutes struct {
	commonAPI zomato.CommonAPI
}

func newZomatoRoutes(commonAPI zomato.CommonAPI) *zomatoRoutes {
	return &zomatoRoutes{commonAPI: commonAPI}
}

func (z *zomatoRoutes) GetZomatoCategoriesHandler(ctx iris.Context) {
	categories, err := z.commonAPI.Categories()
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

func (z *zomatoRoutes) GetZomatoCitiesHandler(ctx iris.Context) {
	//categories, err := z.commonAPI.Categories()
	//if err != nil {
	//	ctx.StatusCode(iris.StatusInternalServerError)
	//	ctx.JSON(
	//		ghttp.HttpError{
	//			Message: fmt.Sprintf("system error, unable to retrieve categories, %s", err.Error()),
	//			Status:  http.StatusInternalServerError,
	//		},
	//	)
	//	return
	//}
	//
	//ctx.StatusCode(iris.StatusOK)
	//ctx.JSON(categories)
}
