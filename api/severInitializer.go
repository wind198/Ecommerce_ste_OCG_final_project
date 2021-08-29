package api

import (
	"database/sql"
	"log"

	"example.com/database/todb"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type dbHandler struct {
	db *sql.DB
}

type Product struct {
	ProductID          int     `json:"productID"`
	ProductName        string  `json:"productName"`
	ProductPrice       float64 `json:"productPrice"`
	ProductImageSource string  `json:"productImageSource"`
}
type User struct {
	UserID       int     `json:"userID"`
	UserEmail    string  `json:"userEmail"`
	UserPassword string  `json:"userPassword"`
	UserBudget   float64 `json:"userBudget"`
}

type ProductPerCategory struct {
	CategoryID   int       `json:"categoryID"`
	CategoryName string    `json:"categoryName"`
	ProductArray []Product `json:"productArray"`
}
type ResponseToLogin struct {
	UserEmail string `json:"userEmail"`
	JWTtoken  string `json:"jwtToken"`
}
type userToCheck struct {
	UserEmail    string
	UserPassword string
}

var newDbHanlder = dbHandler{
	todb.ConnectDB(),
}

func InitiallizeServer() {
	router := mux.NewRouter().StrictSlash(true)
	generalSubrouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	generalSubrouter.HandleFunc("/home", homeHandler).Methods("GET")
	generalSubrouter.HandleFunc("/login", loginHandler).Methods("POST")
	generalSubrouter.HandleFunc("/register", registerHandler).Methods("POST")

	personalRoutes := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)
	personalRoutes.HandleFunc("/user", userHandler).Methods("GET")

	router.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(authenticationMetalware),
		negroni.Wrap(personalRoutes),
	))
	n := negroni.New()
	n.UseHandler(router)
	log.Println("running server")
	n.Run(":5000")

}
