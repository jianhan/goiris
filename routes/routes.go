package routes

import (
	"github.com/jianhan/goiris/bootstrap"
	"github.com/jianhan/goiris/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	b.Get("/", GetIndexHandler)
	b.UseGlobal(middleware.Auth)
}
