package googleplace

import (
	"errors"
	"os"
	"sync"

	"googlemaps.github.io/maps"
)

var (
	instance *maps.Client
	once     sync.Once
	err      error
)

func GetClient() (*maps.Client, error) {
	if os.Getenv("GOOGLE_MAP_API_KEY") == "" {
		return nil, errors.New("api key not setup")
	}
	once.Do(func() {
		instance, err = maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAP_API_KEY")))
	})

	if err != nil {
		return nil, err
	}

	return instance, nil
}
