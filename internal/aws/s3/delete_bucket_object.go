package s3

import (
	"pokeShowcase-api/configs"

	"github.com/aws/aws-sdk-go/aws"
	awsS3 "github.com/aws/aws-sdk-go/service/s3"
)

func DeleteBucketObject(objectKey string) error {
	c := GetS3Connection()

	_, err := c.DeleteObject(
		&awsS3.DeleteObjectInput{
			Bucket: aws.String(configs.Envs.S3_BUCKET_NAME),
			Key:    aws.String(objectKey),
		},
	)
	return err
}
