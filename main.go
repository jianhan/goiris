package main

import (
	"fmt"
	"github.com/jianhan/goiris/bootstrap"
	"github.com/jianhan/goiris/middleware"
	"github.com/jianhan/goiris/routes"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
)

func newApp() *bootstrap.Bootstrapper {
	envs, err := bootstrap.EnvConfigs()
	if err != nil {
		panic(err)
	}
	logrus.Info(envs)

	return bootstrap.New(envs, middleware.AppHeadersConfigure, routes.Configure).Bootstrap()
}

func main() {
	newApp().Listen(
		fmt.Sprintf(":%d", app.Env.Port),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)

}
