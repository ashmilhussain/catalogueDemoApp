package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"io/ioutil"
	"os"
	"strconv"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

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

func (server *Server) Initialize(DbPort, DbHost string) {

	var err error

	Dbdriver := "postgres"
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, "postgresadmin", "postgresdb", "admin123")
	server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	err = server.DB.AutoMigrate(&Product{})

}

func (server *Server) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go , You are visiting Directory :/ %q", html.EscapeString(r.URL.Path))
}

func (server *Server) ProductList(w http.ResponseWriter, r *http.Request) {

	var products []Product
	server.DB.Find(&products)

	for i, product := range products {
		products[i].Price = ConvertCurrenct(product.Price)
	}

	json.NewEncoder(w).Encode(products)
}

func ConvertCurrenct(amount int) int {

	amountstr := strconv.Itoa(amount)
	currencySHost := os.Getenv("CS_HOST")
	currencySPort := os.Getenv("CS_PORT")

	requestURL := "http://" + currencySHost + ":" + currencySPort + "/convert/" + amountstr
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
