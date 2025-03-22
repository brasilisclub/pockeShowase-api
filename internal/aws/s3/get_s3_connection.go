package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Connection *s3.S3

func GetS3Connection(sess *session.Session) *s3.S3 {
	return s3.New(sess, &aws.Config{
		Endpoint:         aws.String("https://s3.amazonaws.com"), // Use o correto para seu caso
		S3ForcePathStyle: aws.Bool(true),
	})
}
