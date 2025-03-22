package s3

import (
	"pokeShowcase-api/internal/aws"

	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Connection *s3.S3

func GetS3Connection() *s3.S3 {
	ses := aws.GetSession()
	if s3Connection == nil {
		s3Connection = newS3Connection(ses)
	}
	return s3Connection
}
