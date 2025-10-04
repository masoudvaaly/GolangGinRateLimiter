package service

import (
	_ "github.com/go-sql-driver/mysql"
	Config "ratelimiter/config"
	"ratelimiter/models"
)

func SavePayment(charge *models.Charge) (err error) {
	if err = Config.GormDB.Create(charge).Error; err != nil {
		return err
	}
	return nil

}
