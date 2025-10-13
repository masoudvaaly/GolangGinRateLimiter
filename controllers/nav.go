package controllers

import (
	"fmt"
	"ratelimiter/config"
)

func GetNAVs(query string) (bool, error) {
	rows, err := config.DB.Query(query)
	if err != nil {
		return false, fmt.Errorf("failed %w", err)
	}
	defer rows.Close()

	return true, nil
}
