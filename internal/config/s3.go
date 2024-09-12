package config

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var globalS3 *s3.S3

func initS3() error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(spaceRegion),
		Endpoint:    aws.String(spaceURL),
		Credentials: credentials.NewStaticCredentials(spaceKey, spaceSecret, ""),
	})
	if err != nil {
		return errors.New("failed to connect to s3: " + err.Error())
	}

	globalS3 = s3.New(sess)
	return nil
}

func GetS3() *s3.S3 {
	return globalS3
}
