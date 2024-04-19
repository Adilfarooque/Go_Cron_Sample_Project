package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func myFunction() {
	fmt.Println("Executing myFunction at:", time.Now())

}

func main() {
	c := cron.New()

	_, err := c.AddFunc("*/5 * * * *", myFunction)
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	c.Start()

	select {}
}
