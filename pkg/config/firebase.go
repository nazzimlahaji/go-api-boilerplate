package config

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	AuthClient *auth.Client
	MsgClient  *messaging.Client
}

func FirebaseConfig() (*FirebaseClient, error) {
	// Initialize Firebase Admin SDK
	opt := option.WithCredentialsFile("./serviceAccountKey.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	// Create Firebase Auth client
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	msgClient, err := app.Messaging(context.Background())
	if err != nil {
		return nil, err
	}

	return &FirebaseClient{
		AuthClient: authClient,
		MsgClient:  msgClient,
	}, nil
}
