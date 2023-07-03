package model

import (
	"time"
)

type Customer struct {
	Id          string
	Name        string
	Email       string
	PlanId      string
	AccessKey   string
	SecretKey   string
	CreatedTime time.Time
	UpdatedTime time.Time
}
