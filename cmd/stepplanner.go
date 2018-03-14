package main

import (
	"flag"
	"fmt"
	"github.com/alltom/stepplanner/calc"
	"os"
	"time"
)

var (
	goal          = flag.Int("goal", -1, "Number of steps you want to take today")
	wakeTimeFlag  = flag.String("wakeTime", "5:00am", "Time you woke up today")
	sleepTimeFlag = flag.String("sleepTime", "9:00pm", "Time you plan to go to sleep today")
)

const (
	timeFormat = "3:04pm"
)

func main() {
	flag.Parse()

	if *goal < 0 {
		fmt.Fprintf(os.Stderr, "invalid step goal: %d\n", *goal)
		flag.PrintDefaults()
		os.Exit(1)
	}

	wakeTime, err := time.Parse(timeFormat, *wakeTimeFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid time format: %s\n", *wakeTimeFlag)
		flag.PrintDefaults()
		os.Exit(1)
	}

	sleepTime, err := time.Parse(timeFormat, *sleepTimeFlag)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid time format: %s\n", *sleepTimeFlag)
		flag.PrintDefaults()
		os.Exit(1)
	}

	now := time.Now()
	dayStart := time.Date(now.Year(), now.Month(), now.Day(), wakeTime.Hour(), wakeTime.Minute(), 0, 0, now.Location())
	dayEnd := time.Date(now.Year(), now.Month(), now.Day(), sleepTime.Hour(), sleepTime.Minute(), 0, 0, now.Location())

	fmt.Printf("You should have %d steps\n", calc.GetSteps(*goal, now, dayStart, dayEnd))
}
