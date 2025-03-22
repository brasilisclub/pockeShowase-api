package aws

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

var ses *session.Session

func GetSession() *session.Session {

	if ses == nil {
		ses = createNewSession()
	}
	return ses

}
