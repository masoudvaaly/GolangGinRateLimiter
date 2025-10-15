package chains

import (
	"github.com/sirupsen/logrus"
	"ratelimiter/controllers"
	"ratelimiter/util"
	"time"
)

type SubordinateRemover struct {
	next Department
}

func (c *SubordinateRemover) Execute(r *Request) {
	logrus.Info("SubordinateRemover started")

	date := "2025-10-13"

	fundId := 1
	isFirstDate := false
	now := time.Now()
	yesterday := now.Add(-24 * time.Hour)

	if SEMAT_CHECK {
		parseDate, err := util.ParseDate(date)
		if err == nil {
			compare := parseDate.Compare(now)
			logrus.WithFields(logrus.Fields{
				"now":         now,
				"compare":     compare,
				"fundId":      fundId,
				"isFirstDate": isFirstDate,
			}).Info()
		} else {
			logrus.Error("date compare err")
		}
	}

	if hasNavInSpecificDay(yesterday) {
		logrus.Error("has nav in yesterday")
	} else if hasNavInSpecificDay(now) {
		logrus.Error("has nav in today")
	} else {
		logrus.Warn("err")
	}

	//
	if !isInFiscalYear(now) {
		logrus.Error("not in current year")
	}

	vouchers, _ := controllers.GetVouchersByType(controllers.OPENING_YEAR)
	if len(vouchers) < 0 {
		logrus.Error("opening voucher not found")
	}
}

func (c *SubordinateRemover) SetNext(next Department) {
	c.next = next
}
