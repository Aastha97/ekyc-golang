package model

import (
	"time"
)

type Customer struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PlanId      uint      `json:"plan" gorm:"foriegn_key"`
	AccessKey   string    `json:"access_key"`
	SecretKey   string    `json:"secret_key"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
