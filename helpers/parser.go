package helpers

import (
	"fmt"
	"strconv"

	"github.com/goofynugtz/TaskTrackr-Scheduler/dao"
	"github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func ParseWorkdayEntryToOrganizationMailReport(workday *[]models.Workday) (map[string][]models.UserReport) {
	// organization_id is mapped to array of user progress for the day.
	var organization map[string][]models.UserReport = make(map[string][]models.UserReport)
	// If workday entry does not exists. The mail will not be fired obvs.
	for _, entry := range *workday {
		user, _ := dao.GetUserById(entry.UserId)
		var report models.UserReport
		report.UserId = user.ID
		report.UserEmail = *user.Email
		report.Name = fmt.Sprintln(user.FirstName, " ", user.LastName)
		
		for project_id, timestring := range entry.LastUpdate {
			time, _ := strconv.Atoi(timestring)
			var details models.Details
			details.Name = *user.Organization.Projects[project_id].ProjectName
			
			details.TimeSpent.Hours = (time / 3600000) % 24
			details.TimeSpent.Minutes = (time / 60000) % 60
			details.TimeSpent.Seconds = (time / 1000) % 60
			report.Projects[project_id] = details
		}
		organization[user.Organization.ID] = append(organization[user.Organization.ID], report)
	}
	fmt.Println(len(organization))
	return organization
}
