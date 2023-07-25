package dao

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func GetUserById(uid string) (*models.User, error) {
	var user models.User
	params := &dynamodb.GetItemInput{
		TableName: aws.String(UsersTable),
		Key: map[string]*dynamodb.AttributeValue{
			"_id": {S: aws.String(uid)},
		},
	}
	resp, err := DynamoDB.GetItem(params)
	if err != nil {
		return nil, err
	}
	if resp.Item == nil {
		return nil, errors.New("_id not found")
	}
	err = dynamodbattribute.UnmarshalMap(resp.Item, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetAllAdministrators(organization_id string) (*[]models.User, error) {
	var administrators []models.User
	input := &dynamodb.ScanInput{
		TableName:        aws.String(UsersTable),
		FilterExpression: aws.String("#organization.#id = :organization_id and #ut = :type"),
		ExpressionAttributeNames: map[string]*string{
			"#organization": aws.String("organization"),
			"#id":           aws.String("_id"),
			"#ut":           aws.String("user_type"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":organization_id": {S: aws.String(organization_id)},
			":type":            {S: aws.String("ADMIN")},
		},
	}
	resp, err := DynamoDB.Scan(input)
	if err != nil {
		return nil, err
	}
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &administrators)
	if err != nil {
		return nil, err
	}
	return &administrators, nil
}
