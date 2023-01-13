package repositorio

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func StartDynamo() (db *dynamodb.DynamoDB) {
	var server = "http://localhost:8000"
	sess := session.Must(session.NewSession())
	db = dynamodb.New(sess, &aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "default"),
		Endpoint:    &server,
	})
	fmt.Printf("Start dynamo %v\n", db)
	return db
}
