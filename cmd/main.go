package main

import (
	"database/sql"
	"ekyc/model"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/brianvoe/gofakeit"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
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

	if !validateEmail(requestBody.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "invalid email"})
		return
	}
	var plan model.Plan
	plan, _ = getPlanByName(requestBody.Plan)
	customer := model.Customer{
		Name:      requestBody.Name,
		Email:     requestBody.Email,
		PlanId:    plan.ID,
		AccessKey: gofakeit.UUID()[:10],
		SecretKey: gofakeit.UUID()[:20],
	}
	//Store customer details in the database
	err := insertCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorMessage": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accessKey": customer.AccessKey,
		"secretKey": customer.SecretKey,
	})
}

func validateEmail(email string) bool {
	regexPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(regexPattern)

	return regex.MatchString(email)
}

func validatePlan(planName string) bool {
	_, validPlan := getPlanByName(planName)
	if validPlan != nil {
		return false
	}
	return true
}

func insertCustomer(customer model.Customer) error {
	statement := `INSERT INTO customer (name, email_id, plan_id, access_key, secret_key) VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(statement, customer.Name, customer.Email, customer.PlanId, customer.AccessKey, customer.SecretKey)
	if err != nil {
		return err
	}
	return nil
}

func getPlanByName(planName string) (model.Plan, error) {
	var plan model.Plan
	query := "SELECT * FROM plan WHERE plan_name = $1"
	err := db.QueryRow(query, planName).Scan(&plan.ID, &plan.PlanName, &plan.DailyBaseCost, &plan.ApiPricing, &plan.StoragePrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Plan{}, fmt.Errorf("plan not found")
		}
		return model.Plan{}, err
	}

	return plan, nil
}
