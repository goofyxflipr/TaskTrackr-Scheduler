package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goofynugtz/TaskTrackr-Scheduler/scheduler"
	"github.com/robfig/cron/v3"
)

func init() {
	os.Setenv("TZ", "Africa/Cairo")
}

func main() {
	router := gin.Default()
	// scheduler.CollectDataForToday()
	c := cron.New()
	if _, err := c.AddFunc("0 17 * * 1-5", scheduler.CollectDataForToday); err != nil {
		fmt.Println("error setting up cronjob: ", err)
		return
	}
	c.Start()
	router.Run(":8080")
}
