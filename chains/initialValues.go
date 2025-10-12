package chains

import (
	"fmt"
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
	//yesterday := now.Add(-24 * time.Hour)

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
}

func (c *InitialValues) SetNext(next Department) {
	c.next = next
}
