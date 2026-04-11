package platform

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
}

func NewS3Client() (S3Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithCredentialsProvider(
			aws.NewCredentialsCache(
				credentials.StaticCredentialsProvider{
					Value: aws.Credentials{
						AccessKeyID:     "test",
						SecretAccessKey: "test",
						SessionToken:    "test",
					},
				},
			),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(opts *s3.Options) {
		opts.BaseEndpoint = aws.String("http://localhost:4566")
		opts.UsePathStyle = true
	})

	out, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		panic(err)
	}

	for _, bucket := range out.Buckets {
		log.Printf(*bucket.Name)
	}

	return S3Client{
		Client: client,
	}, nil
}
