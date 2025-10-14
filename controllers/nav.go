package controllers

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"ratelimiter/config"
)

func GetNAVs(query string) (*sql.Rows, error) {
	logrus.WithFields(logrus.Fields{
		"nav query": query,
	}).Info()
	rows, err := config.DB.Query(query)
	if err != nil {
		return rows, fmt.Errorf("failed %w", err)
	}
	defer rows.Close()

	logrus.WithFields(logrus.Fields{
		"nav rows": rows,
	}).Info()
	return rows, nil
}
