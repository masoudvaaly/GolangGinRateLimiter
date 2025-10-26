package chains

import (
	"fmt"
	"ratelimiter/controllers"
	"ratelimiter/util"
	"time"

	"github.com/sirupsen/logrus"
)

type PreControls struct {
	next Department
}

func (c *PreControls) Execute(r *Request) {
	logrus.Info("PreControls started")

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

	c.next.Execute(r)
}

func isInFiscalYear(now time.Time) bool {
	fiscalYear, err := controllers.GetCurrentFiscalYear()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"fiscal year error ": err,
		}).Warn()
		return false
	}

	logrus.WithFields(logrus.Fields{
		"fiscalYear": fiscalYear.YearName + "-" + fiscalYear.StartDate.String() + "-" + fiscalYear.EndDate.String(),
	}).Info()

	return now.Compare(fiscalYear.StartDate) < 0 || now.Compare(fiscalYear.EndDate) > 0
}

func hasNavInSpecificDay(calcDate time.Time) bool {
	query := fmt.Sprintf("select count(*) from nav n where n.calcDate = %s and light = 1", calcDate.Format("2006-01-02"))
	result, err := controllers.GetNAVs(query)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err nav count": err,
		}).Warn()
		return false
	}

	var count int
	err = result.Scan(&count)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err nav count": err,
		}).Warn()
		return false
	}

	logrus.WithFields(logrus.Fields{
		"nav count": count,
	}).Info()
	return count > 0
}

func (c *PreControls) SetNext(next Department) {
	c.next = next
}
