package main

import (
	"github.com/jianhan/goiris/bootstrap"
	"github.com/jianhan/goiris/middleware/identity"
	"github.com/jianhan/goiris/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8888")
}
