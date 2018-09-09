package bootstrap

import (
	"time"

	"github.com/gorilla/securecookie"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris/websocket"
)

// Configurator is functional options for configurations of bootstrap
type Configurator func(*bootstrapper)

type Bootstrapper interface {
	Bootstrap() Bootstrapper
	Configure(cs ...Configurator)
	Listen(addr string, cfgs ...iris.Configurator)
}

// bootstrapper setup app.
type bootstrapper struct {
	*iris.Application
	AppName       string
	AppOwner      string
	AppOwnerEmail string
	AppSpawnDate  time.Time
	Sessions      *sessions.Sessions
	Env           *Env
}

// New returns a new bootstrapper.
func New(env *Env, cfgs ...Configurator) Bootstrapper {
	b := &bootstrapper{
		AppName:       env.AppName,
		AppOwner:      env.AppOwner,
		AppOwnerEmail: env.AppOwnerEmail,
		AppSpawnDate:  time.Now(),
		Application:   iris.New(),
		Env:           env,
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// setupViews loads the templates.
func (b *bootstrapper) setupViews(viewsDir string) {
	b.RegisterView(iris.HTML(viewsDir, ".html").Layout("shared/layout.html"))
}

// setupSessions initializes the sessions, optionally.
func (b *bootstrapper) setupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

// setupWebsockets prepares the websocket server.
func (b *bootstrapper) setupWebsockets(endpoint string, onConnection websocket.ConnectionFunc) {
	ws := websocket.New(websocket.Config{})
	ws.OnConnection(onConnection)

	b.Get(endpoint, ws.Handler())
	b.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

// setupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
func (b *bootstrapper) setupErrorHandlers() {
	b.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
// Returns itself.
func (b *bootstrapper) Bootstrap() Bootstrapper {
	b.setupViews("./views")
	b.setupSessions(24*time.Hour,
		[]byte(b.Env.CookieHashKey),
		[]byte(b.Env.CookieBlockKey),
	)
	b.setupErrorHandlers()

	// static files
	b.Favicon(StaticAssets + Favicon)
	b.StaticWeb(StaticAssets[1:len(StaticAssets)-1], StaticAssets)

	// middleware, after static files
	b.Use(recover.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
func (b *bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
