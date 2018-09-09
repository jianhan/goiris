package middleware

import (
	"time"

	"github.com/kataras/iris"

	"github.com/jianhan/goiris/bootstrap"
)

// NewAppHeaders a new handler which adds some headers and view data
// describing the application, i.e the owner, the startup time.
func NewAppHeaders(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		// response headers
		ctx.Header("App-Name", b.AppName)
		ctx.Header("App-Owner", b.AppOwner)
		ctx.Header("App-Since", time.Since(b.AppSpawnDate).String())
		ctx.Header("Server", b.Env.Address())
		// view data if ctx.View or c.Tmpl = "$page.html" will be called next.
		ctx.ViewData("AppName", b.AppName)
		ctx.ViewData("AppOwner", b.AppOwner)
		ctx.Next()
	}
}

// AppHeadersConfigure creates a new identity middleware and registers that to the app.
func AppHeadersConfigure(b *bootstrap.Bootstrapper) {
	h := NewAppHeaders(b)
	b.UseGlobal(h)
}
