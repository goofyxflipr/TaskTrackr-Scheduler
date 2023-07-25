package dao

import (
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/goofynugtz/TaskTrackr-Scheduler/connection"
)


// Tables
var UsersTable = os.Getenv("USERS_TABLE")
var UsersEmailIndex = os.Getenv("SECONDARY_USER_INDEX_NAME")
var WorkdayTable = os.Getenv("WORKDAY_TABLE")
var WorkdayUserDateIndex = os.Getenv("SECONDARY_WORKDAY_INDEX_NAME")

var HackathonsTable = os.Getenv("HACKATHONS_TABLE")

// var secondary_hackathons_index = os.Getenv("SECONDARY_HACKATHONS_INDEX_NAME")

// Buckets
var Bucket = os.Getenv("BUCKET_NAME")

// Connections
var DynamoDB = dynamodb.New(connection.GetSession())
var S3Bucket = s3.New(connection.GetSession())
var SES = ses.New(connection.GetSession())
