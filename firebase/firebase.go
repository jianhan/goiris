package firebase

import (
	"fmt"

	"sync"

	firebaseGo "firebase.google.com/go"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var (
	instance *firebaseGo.App
	once     sync.Once
	err      error
)

func NewFirebaseApp() (*firebaseGo.App, error) {
	once.Do(func() {
		opt := option.WithCredentialsFile("./firebase_sdk_key.json")
		if instance, err = firebaseGo.NewApp(context.Background(), nil, opt); err != nil {
			logrus.Fatal(fmt.Errorf("unable to setup firebase app"), err)
			panic(err)
		}
	})

	return instance, nil
}
