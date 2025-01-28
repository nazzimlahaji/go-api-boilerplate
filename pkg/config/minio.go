package config

import (
	"context"
	"errors"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client *minio.Client
	Bucket string
}

func MinioConfig(endpoint, accessKey, secretKey, bucket string, ssl bool) (*MinioClient, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	// Set up the bucket
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	} else if !exists {
		return nil, errors.New("Bucket does not exist")
	}

	return &MinioClient{
		Client: minioClient,
		Bucket: bucket,
	}, nil
}
