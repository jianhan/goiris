package zomato

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/leebenson/conform"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type base struct {
	apiBaseURL string
}

type CitiesRequest struct {
	Q       string   `conform:"trim" json:"q"`
	Lat     string   `conform:"trim" json:"lat"`
	Lon     string   `conform:"trim" json:"lon"`
	CityIDs []string `json:"city_ids"`
	Count   string   `conform:"trim,num" json:"count"`
}

func (c *CitiesRequest) toQueryString() (string, error) {
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
	if len(c.CityIDs) > 0 {
		logrus.Info(c.CityIDs)
		cityIDSlice := []string{}
		for _, v := range c.CityIDs {
			cityIDSlice = append(cityIDSlice, v)
		}
		cityIDStr := strings.Join(cityIDSlice, ",")
		parameters.Add("city_ids", cityIDStr)
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

type CommonAPI interface {
	Categories() ([]*Category, error)
	Cities(request *CitiesRequest) ([]*City, error)
}

type commonAPI struct {
	base
}

func NewCommonAPI() CommonAPI {
	return &commonAPI{base: base{apiBaseURL: os.Getenv("ZOMATO_API_URL")}}
}

func (c *commonAPI) Categories() ([]*Category, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/categories", os.Getenv("ZOMATO_API_URL")), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", os.Getenv("ZOMATO_API_KEY"))
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

	categoryResponse := CategoryResponse{}
	if err := json.Unmarshal(body, &categoryResponse); err != nil {
		return nil, err
	}

	categories := []*Category{}
	for _, v := range categoryResponse.Categories {
		categories = append(categories, &Category{ID: v.Categories.ID, Name: v.Categories.Name})
	}

	return categories, nil
}

func (c *commonAPI) Cities(request *CitiesRequest) ([]*City, error) {
	queryString, err := request.toQueryString()
	if err != nil {
		return nil, err
	}

	var apiUrl *url.URL
	apiUrl, err = url.Parse(os.Getenv("ZOMATO_API_URL"))
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

	req.Header.Add("user-key", os.Getenv("ZOMATO_API_KEY"))
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
