package database

// This is going to be the logic that directly communicates with
// our database

import (
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	TABLE_NAME = "userTable2"
)

type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

func NewDynamoDB() *DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return &DynamoDBClient{
		databaseStore: db,
	}
}

// When we want to "register" a new user, we actually are:
// - checking if a user with this username already exists in our DB
// - if they do, we return an error
// - if they DONT, that is when we INSERT the user into the DB

func (u *DynamoDBClient) DoesUserExist(username string) (bool, error) {
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

func (u *DynamoDBClient) InsertUser(user types.RegisterUser) error {
	// we are assembling our item
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.Password),
			},
		},
	}

	// we want to actually insert the item
	_, err := u.databaseStore.PutItem(item)
	if err != nil {
		return err
	}

	return nil
}
