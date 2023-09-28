package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/ashmilhussain/catalogueDemoApp/routers"
)

func main() {

	var server = routers.Server{}
	godotenv.Load()
	server.InitializeRoutes()
	server.Handler.Initialize(os.Getenv("DB_PORT"), os.Getenv("DB_HOST"))
	server.Run(":8080")
}
