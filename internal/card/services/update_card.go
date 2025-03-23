package services

import (
	"bytes"
	"mime/multipart"
	s3Connector "pokeShowcase-api/internal/aws/s3"
	"pokeShowcase-api/internal/card"
	"pokeShowcase-api/internal/card/services/helpers"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/sirupsen/logrus"
)

func UpdateCard(id string, crb card.CardRequestBody, file *multipart.FileHeader) (card.Card, error) {

	var dbCard card.Card

	src, err := file.Open()
	if err != nil {
		logrus.Info(1)
		logrus.Error(err.Error())
		return dbCard, err
	}
	defer src.Close()

	dbCard = helpers.NewDbCard(crb)

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(src); err != nil {
		logrus.Info(2)
		logrus.Error(err.Error())
		return dbCard, err
	}

	err = s3Connector.CreateBucketObject(
		aws.String(helpers.GetCardPath(dbCard)),
		bytes.NewReader(buf.Bytes()),
		aws.String(file.Header.Get("Content-Type")),
	)

	if err != nil {
		logrus.Info(3)
		logrus.Error(err.Error())
		return card.Card{}, err
	}

	dbCard, err = UpdateDatabaseCard(id, crb)
	if err != nil {
		return dbCard, err
	}

	return dbCard, err

}
