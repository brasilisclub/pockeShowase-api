package aws

import (
	"pokeShowcase-api/configs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"
)

func createNewSession() *session.Session {

	ses, err := session.NewSession(&aws.Config{
		Region: aws.String(configs.Envs.AWS_REGION),
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil
	}

	return ses
}
