package infra

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

import (
	"context"
	"mime/multipart"
	"time"

	"log"

	// v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type BucketClient struct {
	S3Client      *s3.Client
	PresignClient *s3.PresignClient
}

func NewBucketBasics() *BucketClient {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatal("Couldn't load default configuration. Have you set up your AWS account?:", err)
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	return &BucketClient{
		S3Client:      s3Client,
		PresignClient: s3.NewPresignClient(s3Client),
	}
}

func (basics BucketClient) UploadFile(bucketName string, objectKey string, file *multipart.FileHeader) error {
	content, err := file.Open()
	if err != nil {
		log.Printf("Couldn't read the file %v to %v:%v. Here's why: %v\n",
			objectKey, bucketName, objectKey, err)
	}
	_, err = basics.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   content,
	})
	if err != nil {
		log.Printf("Couldn't upload file %v to %v:%v. Here's why: %v\n",
			objectKey, bucketName, objectKey, err)
	}

	return err
}

func (presigner BucketClient) GetObjectURL(
	bucketName string, objectKey string, lifetimeSecs int64) (string, error) {
	request, err := presigner.PresignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(lifetimeSecs * int64(time.Second))
	})
	if err != nil {
		log.Printf("Couldn't get a presigned request to get %v:%v. Here's why: %v\n",
			bucketName, objectKey, err)
	}
	return request.URL, err
}
