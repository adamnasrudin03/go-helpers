package main

import (
	"fmt"
	"time"

	help "github.com/adamnasrudin03/go-helpers"
)

func main() {
	timeUTC7 := help.NewTimeUTC7()

	now := timeUTC7.Now()
	fmt.Println(now)
	fmt.Println(timeUTC7.StartDate(now))
	fmt.Println(timeUTC7.EndDate(now))

	fmt.Println(timeUTC7.StartDateInString("2024-07-18"))
	fmt.Println(timeUTC7.EndDateInString("2024-07-18"))

	parseUtc7, err := timeUTC7.ParseUTC7(help.FormatDateTime, time.Now().Format(help.FormatDateTime))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(parseUtc7)
}
