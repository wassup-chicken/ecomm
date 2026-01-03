package server

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type Auth interface {
	VerifyIDToken(ctx context.Context, idToken string) error
}

type firebaseApp struct {
	firebase *firebase.App
}

func NewAUth() (Auth, error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE"))
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		return nil, err
	}

	return &firebaseApp{
		firebase: app,
	}, nil
}

func (app *firebaseApp) VerifyIDToken(ctx context.Context, idToken string) error {
	client, err := app.firebase.Auth(ctx)
	if err != nil {
		return err
	}

	_, err = client.VerifyIDToken(ctx, idToken)

	if err != nil {
		return err
	}

	return nil
}
