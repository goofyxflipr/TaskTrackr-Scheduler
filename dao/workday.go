package dao

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func GetAllUserDataWorkingOn(date string) (*[]models.Workday, error) {
	var workdays []models.Workday
	fmt.Println(date)
	params := &dynamodb.ScanInput{
		TableName:        aws.String(WorkdayTable),
		FilterExpression: aws.String("#workdate = :query"),
		ExpressionAttributeNames: map[string]*string{
			"#workdate": aws.String("workdate"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":query": {S: aws.String(date)},
		},
	}
	resp, err := DynamoDB.Scan(params)
	if err != nil {
		return nil, err
	}
	fmt.Println("resp: ", resp)
	fmt.Println("count: ", *resp.Count)
	if err := dynamodbattribute.UnmarshalListOfMaps(resp.Items, &workdays); err != nil {
		return nil, err
	}
	return &workdays, nil
}
