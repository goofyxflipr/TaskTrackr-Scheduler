package dao

import (

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/goofynugtz/TaskTrackr-Scheduler/connection"
)

// Tables
var UsersTable = "traktvt-users"
var UsersEmailIndex = "email-index"
var WorkdayTable = "traktvt-workday"
var WorkdayUserDateIndex = "uid-workdate-index"

var HackathonsTable = "traktvt-hackathons"

// var secondary_hackathons_index = os.Getenv("SECONDARY_HACKATHONS_INDEX_NAME")

// Buckets
var Bucket = "traktvt-monitoring-data"

// Connections
var DynamoDB = dynamodb.New(connection.GetSession())
var S3Bucket = s3.New(connection.GetSession())
var SES = ses.New(connection.GetSession())
