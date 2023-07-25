package connection

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"os"
)

func GetSession() *session.Session {
	godotenv.Load()
	access_key := os.Getenv("ACCESS_KEY")
	secret_key := os.Getenv("SECRET_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			access_key,
			secret_key,
			""),
	})
	if err != nil {
		fmt.Println("Could not initialize session")
	}
	return sess
}
