package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jakoubek/dates"
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

	startDate, _ := dates.GetDateFirstOfMonth(year, month)
	endDate, _ := dates.GetDateLastOfMonth(year, month)

	pattern += "\n"
	for d := startDate; d.After(endDate) == false; d = d.AddDate(0, 0, 1) {
		fmt.Printf(pattern, d.Format("2006-01-02"))
	}

}
