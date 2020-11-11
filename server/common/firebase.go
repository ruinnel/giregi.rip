package common

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewAuthClient() (*auth.Client, error) {
	logger := GetLogger()
	config := GetConfig()
	opt := option.WithCredentialsFile(config.FirebaseAdminJsonPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		logger.Panicf("error initializing firebase app: %v", err)
		return nil, err
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		logger.Panicf("error initializing firebase auth client: %v", err)
		return nil, err
	}
	return client, nil
}
