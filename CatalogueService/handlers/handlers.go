package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"time"

	"io/ioutil"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ItemName  string `gorm:"uniqueIndex"`
	Available bool
	Price     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt
}

type Response struct {
	AmountInINR int `json:"Amount_INR"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go , You are visiting Directory :/ %q", html.EscapeString(r.URL.Path))
}

func ProductList(w http.ResponseWriter, r *http.Request) {

	dsn := "host=localhost user=postgresadmin password=admin123 dbname=postgresdb port=5430"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&Product{})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	var products []Product
	db.Find(&products)

	for i, product := range products {
		products[i].Price = ConvertCurrenct(product.Price)
	}

	json.NewEncoder(w).Encode(products)
}

func ConvertCurrenct(amount int) int {

	amountstr := strconv.Itoa(amount)

	requestURL := "http://localhost:5000/convert/" + amountstr
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
	}
	fmt.Printf("client: response body: %s\n", resBody)
	var responseData Response
	err = json.Unmarshal(resBody, &responseData)
	fmt.Println(responseData)

	return responseData.AmountInINR
}
