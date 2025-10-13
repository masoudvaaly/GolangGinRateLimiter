package controllers

import (
	"database/sql"
	"fmt"
	"ratelimiter/config"
)

func GetNAVs(query string) (*sql.Rows, error) {
	rows, err := config.DB.Query(query)
	if err != nil {
		return rows, fmt.Errorf("failed %w", err)
	}
	defer rows.Close()

	return rows, nil
}
