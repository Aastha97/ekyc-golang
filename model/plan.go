package model

type Plan struct {
	ID            uint    `json:"id"`
	PlanName      string  `json:"plan_name"`
	DailyBaseCost float64 `json:"daily_base_cost"`
	MatchScore    float64 `json:"match_score"`
	StoragePrice  float64 `json:"storage_price"`
}
