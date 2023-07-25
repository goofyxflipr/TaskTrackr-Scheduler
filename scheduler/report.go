package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/goofynugtz/TaskTrackr-Scheduler/dao"
	"github.com/goofynugtz/TaskTrackr-Scheduler/helpers"
	"github.com/goofynugtz/TaskTrackr-Scheduler/models"
)

func CollectDataForToday(){
	yyyy, mm, dd := time.Now().Date()
	today := fmt.Sprintf("%v %v, %v", dd, mm, yyyy)
	allUsersToday, err := dao.GetAllUserDataWorkingOn(today)
	if err != nil {
		log.Fatal(err)
		return
	}
	var reports *[]models.Workday
	for _, user := range *allUsersToday {
		out := helpers.ParseWorkdayEntryToMailReport(&user)
	}
}
