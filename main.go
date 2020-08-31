package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jakoubek/dates"
)

func main() {

	var year, month int
	var pattern string = "* [[%s]] "
	var dateFormat string = "2006-01-02"
	if len(os.Args) >= 3 {
		year, _ = strconv.Atoi(os.Args[1])
		month, _ = strconv.Atoi(os.Args[2])
		if len(os.Args) >= 4 {
			pattern = os.Args[3]
		}
		if len(os.Args) == 5 {
			dateFormat = os.Args[4]
		}
	} else {
		var monthMonth time.Month
		year, monthMonth, _ = time.Now().Date()
		month = int(monthMonth)
	}

	startDate, _ := dates.GetDateFirstOfMonth(year, month)
	endDate, _ := dates.GetDateLastOfMonth(year, month)

	pattern += "\n"
	for d := startDate; d.After(endDate) == false; d = d.AddDate(0, 0, 1) {
		if strings.Count(pattern, "%s") > 1 {
			weekday, _ := dates.GetAbbrWeekdayNameGerman(d.Weekday())
			fmt.Printf(pattern, d.Format(dateFormat), weekday)
		} else {
			fmt.Printf(pattern, d.Format(dateFormat))
		}
	}

}

func AbbrWeekday(weekday int) string {
	var abbrWeekday string
	switch weekday {
	case 0:
		abbrWeekday = "So"
	case 1:
		abbrWeekday = "Mo"
	case 2:
		abbrWeekday = "Di"
	case 3:
		abbrWeekday = "Mi"
	case 4:
		abbrWeekday = "Do"
	case 5:
		abbrWeekday = "Fr"
	case 6:
		abbrWeekday = "Sa"
	}
	return abbrWeekday
}
