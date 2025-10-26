package controllers

import (
	"database/sql"
	"ratelimiter/config"

	"github.com/sirupsen/logrus"
)

func GetNAVs(query string) (*sql.Row, error) {
	logrus.WithFields(logrus.Fields{
		"nav query": query,
	}).Debug()
	rows := config.DB.QueryRow(query)
	return rows, nil
}
