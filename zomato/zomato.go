package zomato

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/leebenson/conform"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

type base struct {
	apiBaseURL string
	apiKey     string
}

var (
	commonAPIInstance CommonAPI
	once              sync.Once
)

type CommonAPI interface {
	Categories() ([]*Category, error)
	Cities(request *CitiesRequest) ([]*City, error)
}

type commonAPI struct {
	base
}

func NewCommonAPI(apiKey, apiBaseURL string) (CommonAPI, error) {
	// simple validation
	if strings.Trim(apiKey, " ") == "" {
		return nil, fmt.Errorf("empty api key, %s", apiKey)
	}

	if strings.Trim(apiBaseURL, " ") == "" {
		return nil, fmt.Errorf("empty api base url, %s", apiKey)
	}

	once.Do(func() {
		commonAPIInstance = &commonAPI{base: base{apiBaseURL: apiBaseURL, apiKey: apiKey}}
	})

	return commonAPIInstance, nil
}

func (c *commonAPI) Categories() ([]*Category, error) {
	// init client
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/categories", c.apiBaseURL), nil)
	if err != nil {
		return nil, err
	}

	// set user key
	req.Header.Add("user-key", c.apiKey)

	// make request
	rsp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	// unmarshal response
	categoryResponse := CategoryResponse{}
	if err := json.Unmarshal(body, &categoryResponse); err != nil {
		return nil, err
	}

	// generate categories
	categories := []*Category{}
	for _, v := range categoryResponse.Categories {
		categories = append(categories, &Category{ID: v.Categories.ID, Name: v.Categories.Name})
	}

	return categories, nil
}

func (c *commonAPI) Cities(request *CitiesRequest) ([]*City, error) {
	queryString, err := request.ToQueryString()
	if err != nil {
		return nil, err
	}

	var apiUrl *url.URL
	apiUrl, err = url.Parse(c.apiBaseURL)
	if err != nil {
		return nil, err
	}
	apiUrl.Path += "/cities"
	apiUrl.RawQuery = queryString
	logrus.Info(apiUrl.String())

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, apiUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", c.apiKey)
	rsp, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	defer rsp.Body.Close()

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	logrus.Info(string(body))

	citiesResponse := CitiesResponse{}
	if err := json.Unmarshal(body, &citiesResponse); err != nil {
		return nil, err
	}

	cities := []*City{}
	for _, v := range citiesResponse.LocationSuggestions {
		cities = append(cities, &v)
	}

	return cities, nil
}

type CitiesRequest struct {
	Q       string `conform:"trim" json:"q" schema:"q"`
	Lat     string `conform:"trim" json:"lat" validate:"required,latitude" schema:"lat"`
	Lon     string `conform:"trim" json:"lon" validate:"required,longitude" schema:"lon"`
	CityIDs string `json:"city_ids" schema:"city_ids"`
	Count   string `conform:"trim,num" json:"count" schema:"city_ids"`
}

func (c *CitiesRequest) ToQueryString() (string, error) {
	if err := conform.Strings(c); err != nil {
		return "", nil
	}

	parameters := url.Values{}
	if c.Q != "" {
		parameters.Add("q", c.Q)
	}
	if c.Lat != "" {
		parameters.Add("lat", c.Lat)
	}
	if c.Lon != "" {
		parameters.Add("lon", c.Lon)
	}
	if c.CityIDs != "" {
		parameters.Add("city_ids", c.CityIDs)
	}
	if c.Count != "" {
		parameters.Add("count", c.Count)
	}
	queryStr := parameters.Encode()
	if queryStr == "" {
		return "", errors.New("empty query for cities")
	}

	return queryStr, nil
}
