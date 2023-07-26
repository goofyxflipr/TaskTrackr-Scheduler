package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/goofynugtz/TaskTrackr-Scheduler/dao"
	"github.com/goofynugtz/TaskTrackr-Scheduler/helpers"
	// "github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func CollectDataForToday() {
	yyyy, mm, dd := time.Now().Date()
	today := fmt.Sprintf("%v %v, %v", dd, mm, yyyy)
	// This has data from all different organizations
	// Also this is different for different timezones
	fmt.Println(today)
	allUsersToday, err := dao.GetAllUserDataWorkingOn(today)
	if err != nil {
		log.Fatal(err)
		return
	}
	// organization_id is mapped to array of user progress for the day.
	organizationWiseReport := helpers.ParseWorkdayEntryToOrganizationMailReport(allUsersToday)
	for organization_id, employee_reports := range organizationWiseReport {
		var adminEmails []*string
		administrators, err := dao.GetAllAdministrators(organization_id)
		if err != nil {
			log.Fatal(err)
			return
		}

		for _, admins := range *administrators {
			adminEmails = append(adminEmails, admins.Email)
		}

		helpers.SendMailToAdmins(adminEmails, employee_reports, dao.SES)
		
		for _, report := range employee_reports {
			helpers.SendMailToUser(report.UserEmail, report.Name, report.Projects, dao.SES)
		}
	}
}
