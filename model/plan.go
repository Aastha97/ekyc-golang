package model

type Plan struct {
	ID            uint    `json:"id" gorm:"primary_key"`
	PlanName      string  `json:"plan_name"`
	DailyBaseCost float64 `json:"daily_base_cost"`
	ApiPricing    float64 `json:"api_pricing"`
	StoragePrice  float64 `json:"storage_price"`
}
