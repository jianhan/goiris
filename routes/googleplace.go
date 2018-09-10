package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/jianhan/goiris/googleplace"
	ghttp "github.com/jianhan/goiris/http"
	"github.com/kataras/iris"
	"github.com/leebenson/conform"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type GoogleSearchRequest struct {
	Name      string `conform:"trim" json:"name"`
	Radius    uint   `json:"radius"`
	Location  string `conform:"trim" validate:"required" json:"location"`
	Keyword   string `conform:"trim" json:"keyword"`
	Language  string `conform:"trim" json:"language"`
	MinPrice  string `conform:"trim" schema:"min_price" json:"min_price"`
	MaxPrice  string `conform:"trim" schema:"max_price" json:"max_price"`
	OpenNow   bool   `schema:"open_now" json:"open_now"`
	RankBy    string `conform:"trim" schema:"rank_by" json:"rank_by"`
	PlaceType string `conform:"trim" schema:"type" conform:"trim" json:"price_type"`
	PageToken string `conform:"trim" schema:"page_token" conform:"trim" json:"page_token"`
}

func (s *GoogleSearchRequest) GenerateNearBySearchRequestOptions() ([]googleplace.NearbySearchRequestOption) {
	options := []googleplace.NearbySearchRequestOption{}
	if s.Name != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.Name(s.Name))
	}

	if s.Radius > 0 {
		options = append(options, googleplace.NearbySearchRequestOptions{}.Raidus(s.Radius))
	}

	if s.Location != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.Location(s.Location))
	}

	if s.Keyword != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.Keyword(s.Keyword))
	}

	if s.Language != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.Language(s.Language))
	}

	if s.MinPrice != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.MinPrice(s.MinPrice))
	}

	if s.MaxPrice != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.MaxPrice(s.MaxPrice))
	}

	options = append(options, googleplace.NearbySearchRequestOptions{}.OpenNow(s.OpenNow))

	if s.RankBy != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.RankBy(s.RankBy))
	}

	if s.PlaceType != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.Type(s.PlaceType))
	}

	if s.PageToken != "" {
		options = append(options, googleplace.NearbySearchRequestOptions{}.PageToken(s.PageToken))
	}

	return options
}

func GetGooglePlaceHandler(ctx iris.Context) {
	// query string to struct
	searchRequest := new(GoogleSearchRequest)
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
					Message: "Validation error",
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
					Message: "Validation error",
					Status:  iris.StatusUnprocessableEntity,
					Data:    validationError,
				},
			)
			return
		}
	}

	// generate search request
	nsReq, rErr := googleplace.NewNearbySearchRequest(searchRequest.GenerateNearBySearchRequestOptions()...)
	if rErr != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("unable generate search options, %s", rErr.Error()),
				Status:  iris.StatusInternalServerError,
			},
		)
		return
	}

	// get client
	client, err := googleplace.GetClient()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("unable to initialize google place client, %s", err.Error()),
				Status:  iris.StatusInternalServerError,
			},
		)
		return
	}

	// call API
	sRsp, sErr := client.NearbySearch(context.Background(), nsReq)
	if sErr != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(
			ghttp.HttpError{
				Message: fmt.Sprintf("unable fetch search results, %s", sErr.Error()),
				Status:  iris.StatusInternalServerError,
			},
		)
		return
	}

	// return response
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(sRsp)
}
