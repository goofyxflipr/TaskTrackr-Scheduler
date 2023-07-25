package helpers

import (
	"fmt"
	"strconv"

	"github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func ParseWorkdayEntryToMailReport(workday *models.Workday) (*models.Report) {
	var report models.Report
	report.UserId = workday.UserId
	for key, value := range workday.LastUpdate {
		time, _ := strconv.Atoi(value)
		var ts models.TimeSpent
		ts.Hours = (time / 3600000) % 24
		ts.Minutes = (time / 60000) % 60
		ts.Seconds = (time / 1000) % 60
		report.Projects[key] = ts
	}
	fmt.Println(report)
	return &report
}
