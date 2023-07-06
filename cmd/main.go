package main

import (
	"ekyc/model"
	"log"
	"net/http"

	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
)

var plans = map[string]model.Plan{
	"basic": {
		PlanName:      "basic",
		DailyBaseCost: 10,
		MatchScore:    10.0,
		StoragePrice:  25.0,
	},
	"advanced": {
		PlanName:      "advanced",
		DailyBaseCost: 15,
		MatchScore:    20.0,
		StoragePrice:  30.0,
	},
	"enterprise": {
		PlanName:      "enterprise",
		DailyBaseCost: 20,
		MatchScore:    35.0,
		StoragePrice:  40.0,
	},
}

//var db *sql.DB

func main() {
	// cfg := config.Load()
	// sqlWriter := writer.NewPostgresWriter(cfg.DBConfig)
	// sqlWriter.WriteSQL()

	router := gin.Default()
	router.POST("/api/v1/signup", handleSignup)
	log.Fatal(router.Run(":8080"))
}

func handleSignup(c *gin.Context) {
	var requestBody struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Plan  string `json:"plan"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		return
	}

	if !validatePlan(requestBody.Plan) {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "invalid plan"})
		return
	}

	customer := model.Customer{
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		PlanId:    gofakeit.UUID(),
		AccessKey: gofakeit.UUID()[:10],
		SecretKey: gofakeit.UUID()[:20],
	}

	// Store customer details in the database
	// err := insertCustomer(customer)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"accessKey": customer.AccessKey,
		"secretKey": customer.SecretKey,
	})
}

func validatePlan(plan string) bool {
	_, validPlan := plans[plan]
	return validPlan
}

// func insertCustomer(customer model.Customer) error {
// 	statement := `INSERT INTO customer (name, email, plan_id, access_key, secret_key) VALUES ($1, $2, $3, $4, $5)`
// 	_, err := db.Exec(statement, customer.Name, customer.Email, customer.PlanId, customer.AccessKey, customer.SecretKey)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
