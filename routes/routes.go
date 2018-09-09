package routes

import (
	"github.com/jianhan/goiris/bootstrap"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.bootstrapper) {
	b.Get("/", GetIndexHandler)
}
