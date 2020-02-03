package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	var year, month int
	var pattern string = "* [[%s]] "
	if len(os.Args) >= 3 {
		year, _ = strconv.Atoi(os.Args[1])
		month, _ = strconv.Atoi(os.Args[2])
		if len(os.Args) == 4 {
			pattern = os.Args[3]
		}
	} else {
		fmt.Println("Year/month is missing.")
		fmt.Println("Usage:")
		fmt.Println("  all_dates_in_month 2020 6")
		fmt.Println("or")
		fmt.Println("  all_dates_in_month 2020 6 \"* %s\"")
		os.Exit(1)
	}

	startDate := date(year, month, 1)
	endDate := date(year, month, 31)

	pattern += "\n"
	for d := startDate; d.After(endDate) == false; d = d.AddDate(0, 0, 1) {
		fmt.Printf(pattern, d.Format("2006-01-02"))
	}

}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
