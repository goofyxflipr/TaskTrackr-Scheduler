package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func mailWorkdayReport() {

}

func main() {
	c := cron.New()
	if _, err := c.AddFunc("0 17 * * 1-5", mailWorkdayReport); err != nil {
		fmt.Println("error setting up cronjob: ", err)
		return
	}
	c.Start()
}
