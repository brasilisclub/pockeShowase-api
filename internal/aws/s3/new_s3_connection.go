package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func newS3Connection(ses *session.Session) *s3.S3 {
	return s3.New(ses, &aws.Config{
		Endpoint:         aws.String("https://s3.amazonaws.com"),
		S3ForcePathStyle: aws.Bool(true),
	})
}
