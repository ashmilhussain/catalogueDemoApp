package routers

import (
	"net/http"

	"fmt"
	"log"

	"github.com/gorilla/mux"

	myHandler "github.com/ashmilhussain/catalogueDemoApp/handlers"
)

type Server struct {
	Handler myHandler.Server
	Router  *mux.Router
}

func (server *Server) InitializeRoutes() {

	server.Router = mux.NewRouter()
	server.addRoutes()

}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func (server *Server) addRoutes() {
	server.Router.HandleFunc("/", server.Handler.Index).Methods("GET")
	server.Router.HandleFunc("/products", server.Handler.ProductList).Methods("GET")
}
