package ddb

import (
	"AwsServerLessCleanCodeArchitecture/entity"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

const (
	table = "user-table"
)

type DynamoDB struct {
	dynamodb dynamodbiface.DynamoDBAPI
}

func NewDynamoDB(dynamodbClint dynamodbiface.DynamoDBAPI) *DynamoDB {
	return &DynamoDB{dynamodb: dynamodbClint}
}

func (d *DynamoDB) GetUser(username string) (details entity.User, err error) {
	result, err := d.dynamodb.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})

	if err != nil {
		return details, fmt.Errorf("failed to unmarshal results with error: %w", err)
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &details)

	if err != nil {
		return details, fmt.Errorf("failed to unmarshal results with error: %w", err)
	}

	return details, nil

}

func (d *DynamoDB) AddUser(user entity.User) error {
	_, err := d.dynamodb.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.Password),
			},
			"address": {
				S: aws.String(user.Address),
			},
			"first_name": {
				S: aws.String(user.FirstName),
			},
			"last_name": {
				S: aws.String(user.LastName),
			},
		},
	})
	if err != nil {
		return fmt.Errorf("failed to insert to dynamodb with error: %v", err)
	}

	return nil
}
