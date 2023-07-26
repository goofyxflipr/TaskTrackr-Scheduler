package helpers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func MailTemplateLayout(children string) string {
	return `
	<div style='position: relative; padding: 5rem; width: 420px; margin: 0 auto; background-color: white;'>
		<div style='display: flex; width: 20rem;'>
			<div style='display: flex; width: 5rem; height: 5rem;'>
				<img src="https://drive.google.com/uc?export=view&id=1vuea6Ujmsmh40uCwEyxUgXa0TUeKiNyN" style="height: 5rem; width: 5rem;"/>
			</div>
			<div style='color: black; font-size: 20px; font-family: Inter; font-weight: 400; margin-left: 1rem; margin-top: .5rem;'>
				<div style='font-size: 30px;'>TaskTrackr</div>
				<div style='font-size: 20px;'>Workplace Assistant</div>
			</div>
		</div>
		` + children + `
	</div>
	`
}

func SendMailToAdmins(admin_emails []*string, reports []models.UserReport, svc *ses.SES) {
	host_email := os.Getenv("HOST_EMAIL")
	fmt.Println(host_email)
	body :=
		`
	<div style='margin-top: 1rem; font-size: 18px; font-family: Inter; font-weight: 400;'>
	`
	fmt.Println("10")
	for _, report := range reports {
		fmt.Println("11")

		body +=
			`
		<div>
			<div>
				<span>` + report.Name + `</span>
				<span style='color: #0D92FF;'>` + report.UserEmail + `</span>
				<span style='color: black'>Joined today at: ` + report.TimeJoined + `</span>
			</div>
		`
		fmt.Println("12")

		for _, project := range report.Projects {
			fmt.Println("13")
			fmt.Println(project)
			body += `
			<div>
				<span>Project: ` + *project.Name + `</span>
				<span>Time Spent: ` + strconv.Itoa(project.TimeSpent.Hours) + `:` + strconv.Itoa(project.TimeSpent.Minutes) + `:` + strconv.Itoa(project.TimeSpent.Seconds) + `</span>
			</div>
			`
		}
		fmt.Println("14")

		body += `
		</div>
		`
	}
	fmt.Println("15")

	body += `
	</div>
	`
	fmt.Println(admin_emails)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: admin_emails,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(MailTemplateLayout(body)),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Admin Employee Report"),
			},
		},
		Source: aws.String(host_email),
	}
	_, err := svc.SendEmail(input)
	if err != nil {
		return
	}
}

func SendMailToUser(email string, name string, user_report map[string]models.Details, svc *ses.SES) {
	host_email := os.Getenv("HOST_EMAIL")
	body :=
		`
	<div style='margin-top: 1rem; font-size: 18px; font-family: Inter; font-weight: 400;'>
	`
	for _, projectDetails := range user_report {
		body += `
		<div>
			<span>Project: ` + *projectDetails.Name + `</span>
			<span>Time Spent: ` + strconv.Itoa(projectDetails.TimeSpent.Hours) + `:` + strconv.Itoa(projectDetails.TimeSpent.Minutes) + `:` + strconv.Itoa(projectDetails.TimeSpent.Seconds) + `</span>
		</div>
		`
	}
	body += `
	</div>
	`
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(MailTemplateLayout(body)),
				},
			},
			Subject: &ses.Content{
				Data: aws.String("Daily Progress Report"),
			},
		},
		Source: aws.String(host_email),
	}
	_, err := svc.SendEmail(input)
	if err != nil {
		return
	}
}
