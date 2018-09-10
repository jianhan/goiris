package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	ghttp "github.com/jianhan/goiris/http"
	"github.com/jianhan/goiris/zomato"
	"github.com/kataras/iris"
	"github.com/leebenson/conform"
	"gopkg.in/go-playground/validator.v9"
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
	// query string to struct
	searchRequest := new(zomato.CitiesRequest)
	schema.NewDecoder().Decode(searchRequest, ctx.Request().URL.Query())
	conform.Strings(&searchRequest)
	_, err := json.Marshal(&searchRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("system error, unable to marshal request, %s", err.Error()),
				Status:  http.StatusBadRequest,
			},
		)
		return
	}

	// validation
	if err = validator.New().Struct(*searchRequest); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(
				ghttp.HttpError{
					Message: fmt.Sprintf("Validation error, %s", err.Error()),
					Status:  http.StatusBadRequest,
				},
			)
			return
		}

		// generate validation error
		validationError := map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			validationError[err.Field()] = err.Tag()
		}

		if len(validationError) > 0 {
			ctx.StatusCode(iris.StatusUnprocessableEntity)
			ctx.JSON(
				ghttp.HttpError{
					Message: fmt.Sprintf("Validation error, %s", err.Error()),
					Status:  iris.StatusUnprocessableEntity,
					Data:    validationError,
				},
			)
			return
		}
	}

	// call API to get cities
	cities, err := z.commonAPI.Cities(searchRequest)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("Unable to retrieve cities, %s", err.Error()),
				Status:  iris.StatusInternalServerError,
			},
		)
		return
	}

	//	return cities
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(cities)

	return
}
