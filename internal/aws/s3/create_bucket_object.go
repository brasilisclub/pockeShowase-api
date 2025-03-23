package s3

import (
	"bytes"
	"pokeShowcase-api/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func CreateBucketObject(key *string, body *bytes.Reader, contentType *string) error {
	s3Connection := GetS3Connection()

	_, err := s3Connection.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(configs.Envs.S3_BUCKET_NAME),
		Key:         key,
		Body:        body,
		ContentType: contentType,
	})

	return err
}
