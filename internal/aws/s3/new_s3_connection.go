package s3

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func newS3Connection(ses *session.Session) *s3.S3 {
	return s3.New(ses)
}
