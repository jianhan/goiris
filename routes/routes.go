package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/jianhan/goiris/bootstrap"
	"github.com/jianhan/goiris/zomato"
	"github.com/kataras/iris/cache"
	"os"
	"time"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	commonAPI, err := zomato.NewCommonAPI(os.Getenv("ZOMATO_API_KEY"), os.Getenv("ZOMATO_API_URL"))
	if err != nil {
		panic(err)
	}

	zomatonRoutes := newZomatoRoutes(commonAPI)

	// setup API routes
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	// setup API routes
	apiV1Routes := b.Party("/api/v1", crs).AllowMethods()
	{
		cacheHandler := cache.Handler(1 * time.Hour)
		// google
		googleRoutes := apiV1Routes.Party("/google")
		{
			googleRoutes.Get("/place", cacheHandler, GetGooglePlaceHandler)
		}

		// zomato
		zomatoRoutes := apiV1Routes.Party("/zomato")
		{
			zomatoRoutes.Get("/categories", cacheHandler, zomatonRoutes.GetZomatoCategoriesHandler)
			zomatoRoutes.Get("/cities", cacheHandler, zomatonRoutes.GetZomatoCitiesHandler)
		}
	}
}
