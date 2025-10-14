package controllers

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"ratelimiter/config"
	"time"
)

type FiscalYear struct {
	ID        int
	YearName  string
	IsActive  bool
	StartDate time.Time
	EndDate   time.Time
}

func GetCurrentFiscalYear() (FiscalYear, error) {
	query := "select * from FiscalYear where isActive = true"
	logrus.WithFields(logrus.Fields{
		"fiscal year query": query,
	}).Info()
	fy := FiscalYear{}
	rows, err := config.DB.Query(query)
	if err != nil {
		return fy, fmt.Errorf("failed %w", err)
	}

	defer rows.Close()
	err = rows.Scan(&fy)
	if err != nil {
		return fy, err
	}
	logrus.WithFields(logrus.Fields{
		"fiscal year": rows,
	}).Info()
	return fy, nil
}
