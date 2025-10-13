package chains

import (
	"fmt"
	"ratelimiter/controllers"
	"ratelimiter/util"
	"time"
)

type InitialValues struct {
	next Department
}

func (c *InitialValues) Execute(r *Request) {
	fmt.Println("process request")
	if r.Status == Pending {
		fmt.Println("process started")
	}

	date := "2025-10-13"

	fundId := 1
	isFirstDate := false
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)

	if SEMAT_CHECK {
		parseDate, err := util.ParseDate(date)
		if err == nil {
			compare := parseDate.Compare(now)
			fmt.Println("now:", now)
			fmt.Println("date compare ", compare, "fundId:", fundId, "isFirstDate:", isFirstDate)
		} else {
			fmt.Println("date compare err")
		}
	}

	if hasNavInSpecificDay(yesterday) {
		fmt.Println("has nav in yesterday")
	} else if hasNavInSpecificDay(now) {
		fmt.Println("has nav in today")
	}

}

func hasNavInSpecificDay(calcDate time.Time) bool {
	query := fmt.Sprintf("select count(*) from nav n where n.calcDate = %s and light = 1", calcDate.Format("YYYY/MM/DD"))
	result, err := controllers.GetNAVs(query)
	if err != nil {
		return false
	}

	var count int
	err = result.Scan(&count)
	if err != nil {
		return false
	}

	return count > 0
}

func (c *InitialValues) SetNext(next Department) {
	c.next = next
}
