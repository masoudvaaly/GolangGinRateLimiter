package controllers

import (
	"fmt"
	"ratelimiter/config"
	"time"

	"github.com/sirupsen/logrus"
)

type FiscalYear struct {
	ID        int
	YearName  string
	IsActive  bool
	StartDate time.Time
	EndDate   time.Time
}

func GetCurrentFiscalYear() (FiscalYear, error) {
	query := "select * from FiscalYear where is_active = true"
	logrus.WithFields(logrus.Fields{
		"fiscal year query": query,
	}).Info()
	fy := FiscalYear{}
	rows, err := config.DB.Query(query)
	if err != nil {
		return fy, fmt.Errorf("failed %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&fy); err != nil {
			logrus.WithFields(logrus.Fields{
				"fiscal year Scan error": rows,
			}).Error()
		}
		fmt.Println(id, name)
	}

	if err := rows.Err(); err != nil {
		logrus.WithFields(logrus.Fields{
			"fiscal year error": rows,
		}).Error()
	}

	return fy, nil
}
