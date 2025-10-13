package controllers

import (
	"database/sql"
	"fmt"
	"ratelimiter/config"
)

func GetNAVs(query string) (*sql.Rows, error) {
	fmt.Println("nav query ", query)
	rows, err := config.DB.Query(query)
	if err != nil {
		return rows, fmt.Errorf("failed %w", err)
	}
	//defer rows.Close()

	fmt.Println("nav rows %s", rows)
	return rows, nil
}
