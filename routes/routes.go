package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/jianhan/goiris/bootstrap"
	"github.com/jianhan/goiris/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", middleware.Auth, GetIndexHandler)

	// setup API routes
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	apiV1Routes := b.Party("/api/v1", crs).AllowMethods()
	{
		googleRoutes := apiV1Routes.Party("/google")
		{
			googleRoutes.Get("/place", GetGooglePlaceHandler)
		}
	}

}
