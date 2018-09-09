package middleware

import (
	"context"
	"github.com/jianhan/goiris/firebase"
	"github.com/kataras/iris"
	"strings"
)

func Auth(ctx iris.Context) {
	idToken := ctx.GetHeader("Authorization")
	if idToken == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"status": iris.StatusUnauthorized, "message": "unauthorized"})
		return
	}

	// get token
	splitToken := strings.Split(idToken, "Bearer")
	idToken = strings.Trim(splitToken[1], " ")
	if idToken == "" {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"status": iris.StatusUnauthorized, "message": "unauthorized, id token is missing"})
		return
	}

	// get firebase app
	firebaseApp, err := firebase.NewFirebaseApp()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"status": iris.StatusInternalServerError, "message": "internal server error, unable to authenticate user"})
		return
	}

	// validate user
	client, err := firebaseApp.Auth(context.Background())
	if err != nil {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"status": iris.StatusUnauthorized, "message": "invalid id token"})
		return
	}

	// get user
	if _, err = client.VerifyIDToken(context.Background(), idToken); err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"status": iris.StatusInternalServerError, "message": "unable to verify token"})
		return
	}

	// passed authentication
	ctx.Next()
}
