package model

import (
	"time"
)

type Customer struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PlanId      string    `json:"plan_id"`
	AccessKey   string    `json:"access_key"`
	SecretKey   string    `json:"secret_key"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}
