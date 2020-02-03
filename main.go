package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
		var monthMonth time.Month
		year, monthMonth, _ = time.Now().Date()
		month = int(monthMonth)
	}

	startDate, _ := dates.GetDateFirstOfMonth(year, month)
	endDate, _ := dates.GetDateLastOfMonth(year, month)

	pattern += "\n"
	for d := startDate; d.After(endDate) == false; d = d.AddDate(0, 0, 1) {
		fmt.Printf(pattern, d.Format("2006-01-02"))
	}

}
