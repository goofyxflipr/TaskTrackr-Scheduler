package dao

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func GetAllUserDataWorkingOn(date string) (*[]models.Workday, error) {
	var workdays []models.Workday
	params := &dynamodb.QueryInput{
		TableName:              aws.String(WorkdayTable),
		IndexName:              aws.String(WorkdayUserDateIndex),
		KeyConditionExpression: aws.String("workdate = :workdate"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":workdate": {S: aws.String(date)},
		},
	}
	resp, err := DynamoDB.Query(params)
	if err != nil {
		return nil, err
	}
	if err := dynamodbattribute.UnmarshalListOfMaps(resp.Items, &workdays); err != nil {
		return nil, err
	}
	return &workdays, nil
}
