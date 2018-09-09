package googleplace

import (
	"fmt"
	ghttp "github.com/jianhan/goiris/http"
	"googlemaps.github.io/maps"
	"net/http"
	"strings"
)

const MaxRadius = 50000

type NearbySearchRequestOption func(*maps.NearbySearchRequest) *ghttp.HttpError

func NewNearbySearchRequest(opts ...NearbySearchRequestOption) (*maps.NearbySearchRequest, *ghttp.HttpError) {
	r := &maps.NearbySearchRequest{}
	for _, opt := range opts {
		if err := opt(r); err != nil {
			return nil, err
		}
	}

	return r, nil
}

type NearbySearchRequestOptions struct {
}

func (n NearbySearchRequestOptions) parsePriceLevel(fieldName, priceLevel string) (maps.PriceLevel, *ghttp.HttpError) {
	switch priceLevel {
	case "0":
		return maps.PriceLevelFree, nil
	case "1":
		return maps.PriceLevelInexpensive, nil
	case "2":
		return maps.PriceLevelModerate, nil
	case "3":
		return maps.PriceLevelExpensive, nil
	case "4":
		return maps.PriceLevelVeryExpensive, nil
	default:
		return "", &ghttp.HttpError{
			Message: fmt.Sprintf("invalid %s", fieldName),
			Status:  http.StatusBadRequest,
		}
	}
}

func (n NearbySearchRequestOptions) Raidus(radius uint) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		if radius < 1 || radius > MaxRadius {
			return &ghttp.HttpError{
				Message: fmt.Sprintf("invalid radius parameter, must be greater than 1 and less than %d", MaxRadius),
				Status:  http.StatusBadRequest,
			}
		}
		args.Radius = radius

		return nil
	}
}

func (n NearbySearchRequestOptions) Location(location string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		latLng, err := maps.ParseLatLng(location)
		if err != nil {
			return &ghttp.HttpError{
				Message: fmt.Sprintf("unable to parse lat and lng, %s", err.Error()),
				Status:  http.StatusBadRequest,
			}
		}
		args.Location = &latLng

		return nil
	}
}

func (n NearbySearchRequestOptions) Keyword(keyword string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		if strings.Trim(keyword, " ") == "" {
			return &ghttp.HttpError{
				Message: "keyword can not be empty",
				Status:  http.StatusBadRequest,
			}
		}
		args.Keyword = keyword

		return nil
	}
}

func (n NearbySearchRequestOptions) Language(language string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		if strings.Trim(language, " ") == "" {
			return &ghttp.HttpError{
				Message: "language can not be empty",
				Status:  http.StatusBadRequest,
			}
		}
		args.Language = language

		return nil
	}
}

func (n NearbySearchRequestOptions) MinPrice(minPrice string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) (err *ghttp.HttpError) {
		if args.MinPrice, err = n.parsePriceLevel("min price", minPrice); err != nil {
			return err
		}

		return nil
	}
}

func (n NearbySearchRequestOptions) MaxPrice(maxPrice string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) (err *ghttp.HttpError) {
		if args.MaxPrice, err = n.parsePriceLevel("max price", maxPrice); err != nil {
			return err
		}

		return nil
	}
}

func (n NearbySearchRequestOptions) Name(name string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		if strings.Trim(name, " ") == "" {
			return &ghttp.HttpError{
				Message: "name can not be empty",
				Status:  http.StatusBadRequest,
			}
		}

		args.Name = name

		return nil
	}
}

func (n NearbySearchRequestOptions) OpenNow(openNow bool) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		args.OpenNow = openNow
		return nil
	}
}

func (n NearbySearchRequestOptions) RankBy(rankBy string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		switch rankBy {
		case "prominence":
			args.RankBy = maps.RankByProminence
		case "distance":
			args.RankBy = maps.RankByDistance
		default:
			return &ghttp.HttpError{
				Message: fmt.Sprintf("unknown rank by: \"%v\"", rankBy),
				Status:  http.StatusBadRequest,
			}
		}

		return nil
	}
}

func (n NearbySearchRequestOptions) Type(placeType string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		t, tErr := maps.ParsePlaceType(placeType)
		if tErr != nil {
			return &ghttp.HttpError{
				Message: fmt.Sprintf("unable to parse place type: %s", tErr.Error()),
				Status:  http.StatusBadRequest,
			}
		}
		args.Type = t

		return nil
	}
}

func (n NearbySearchRequestOptions) PageToken(pageToken string) NearbySearchRequestOption {
	return func(args *maps.NearbySearchRequest) *ghttp.HttpError {
		if strings.Trim(pageToken, " ") == "" {
			return &ghttp.HttpError{
				Message: "page token can not be empty",
				Status:  http.StatusBadRequest,
			}
		}
		args.PageToken = pageToken

		return nil
	}
}
