package main

import (
	"github.com/jianhan/goiris/bootstrap"
	"github.com/jianhan/goiris/middleware/identity"
	"github.com/jianhan/goiris/routes"
	"github.com/sirupsen/logrus"
)

func newApp() *bootstrap.Bootstrapper {
	envs, err := bootstrap.EnvConfigs()
	if err != nil {
		panic(err)
	}
	logrus.Info(envs)

	return bootstrap.New(envs, identity.Configure, routes.Configure).Bootstrap()
}

func main() {
	app := newApp()
	app.Listen(":8888")
}
