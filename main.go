package main

import (
	"github.com/jianhan/goiris/bootstrap"
	"github.com/jianhan/goiris/middleware"
	"github.com/jianhan/goiris/routes"
	"github.com/kataras/iris"
)

// newApp returns a new bootstrapper instance.
func newApp() *bootstrap.Bootstrapper {
	envs, err := bootstrap.EnvConfigs()
	if err != nil {
		panic(err)
	}

	return bootstrap.New(envs, middleware.AppHeadersConfigure, middleware.APIHeadersConfigure, routes.Configure).Bootstrap()
}

func main() {
	newApp().Listen(
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
		iris.WithoutStartupLog,
	)
}
