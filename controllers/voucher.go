package controllers

import (
	"fmt"
	"ratelimiter/config"
	"time"

	"github.com/sirupsen/logrus"
)

type voucherTypeEnum int

const (
	OPENING_YEAR voucherTypeEnum = iota
	CLOSING_YEAR
)

type Voucher struct {
	ID              int
	voucherType     voucherTypeEnum
	createDate      bool
	transactionDate time.Time
	EndDate         time.Time
}

func GetVouchersByType(voucherType voucherTypeEnum) ([]Voucher, error) {
	query := "select * from Vouchers where is_active = true"
	logrus.WithFields(logrus.Fields{
		"fiscal year query": query,
	}).Info()
	vouchers := []Voucher{}
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed %w", err)
	}

	defer rows.Close()
	for rows.Next() {
		var voucher Voucher
		err = rows.Scan(&voucher)
		if err != nil {
			return nil, err
		}
		vouchers = append(vouchers, voucher)
	}
	logrus.WithFields(logrus.Fields{
		"vouchers": vouchers,
	}).Info()
	return vouchers, nil
}
