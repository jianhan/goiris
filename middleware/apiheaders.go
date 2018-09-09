package middleware

import (
	"github.com/jianhan/goiris/bootstrap"
	"github.com/kataras/iris"
)

// newAPIHeaders returns a new API headers middleware
func newAPIHeaders(b *bootstrap.Bootstrapper) iris.Handler {
	return func(ctx iris.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Next()
	}
}

// APIHeadersConfigure creates API headers for app.
func APIHeadersConfigure(b *bootstrap.Bootstrapper) {
	h := newAPIHeaders(b)
	b.UseGlobal(h)
}
