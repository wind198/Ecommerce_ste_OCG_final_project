package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"github.com/tuanlh1908developer/ecommer_site_1/connectdb"
	"github.com/urfave/negroni"
)

type dbHandler struct {
	db *sql.DB
}

type Product struct {
	ID    int     `json:"productid"`
	Name  string  `json:"productname"`
	Price float64 `json:"productprice"`
}
type User struct {
	ID       int    `json:"userid"`
	Email    string `json:"useremail"`
	Password string `json:"userpassword"`
}

var newDbHanlder = dbHandler{
	connectdb.ConnectDB(),
}

func InitiallizeServer() {
	r := mux.NewRouter()
	s := r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/productBrowse", productBrowseHandler).Methods("GET")
	s.HandleFunc("/login", loginHandler).Methods("POST")
	s.HandleFunc("/register", registerHandler).Methods("POST")
	n := negroni.New()
	n.UseHandler(r)
	n.Run(":8080")

}

func productBrowseHandler(w http.ResponseWriter, r *http.Request) {
	// todo
	//initialize a query to find all products from database
	//loop through the rows result, scan it to a product struct
	//apend the product to the product slice
	//write response Header
	//endcode the slice into a json and write it to the response
	productList := make([]Product, 0)

	q := `select * from product`
	db := newDbHanlder.db
	rows, err := db.Query(q)
	if err != nil {
		log.Println("Error during querying product table", err)
	}

	for rows.Next() {
		var aProduct Product
		if err := rows.Scan(&aProduct.ID, &aProduct.Name, &aProduct.Price); err != nil {
			log.Println("Error during Scanning:", err)
		}
		productList = append(productList, aProduct)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		log.Println("Error during query and iteration,", err)
	}
	productListJson, err := json.Marshal(productList)
	if err != nil {
		log.Println("Error maping slice to json, ", err)

	}
	w.WriteHeader(http.StatusFound)
	w.Write(productListJson)

}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	//todo
	//get user data from request body
	//decode it into a user struct
	//initialize a query to fiter from database a user with such username and password
	//if found return that user struct to response
	//else return an empty struct to response
	db := newDbHanlder.db
	var userToCheck User
	err := json.NewDecoder(r.Body).Decode(&userToCheck)
	if err != nil {
		log.Println("Error during decoding request body", err)
	}
	q := `select * from user where userEmail = ? and userPassword=?`
	row := db.QueryRow(q, &userToCheck.Email, &userToCheck.Password)
	var userOK User
	err = row.Scan(&userOK.ID, &userOK.Email, &userOK.Password)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "No such user info")
		return
	}
	if err != nil {
		log.Println("Error during query and scan", err)
		return
	}
	userOKJson, err := json.Marshal(userOK)
	if err != nil {
		log.Println("Error during mapping json")
		return
	}
	w.WriteHeader(http.StatusFound)
	w.Write(userOKJson)

}
func registerHandler(w http.ResponseWriter, r *http.Request) {
	//todo
	//get user data from request body
	//create a userdata struct and decode request body into it
	//connect to database and check if a user with such email already exist
	//if not exist add it to db and anwser to response
	db := newDbHanlder.db
	var userToCheck User
	err := json.NewDecoder(r.Body).Decode(&userToCheck)
	if err != nil {
		log.Println("Error during decoding request body", err)
		return
	}
	var validEmail = regexp.MustCompile(`(\w+\d*)@(\w+)\.(net|com|org)`)
	m := validEmail.FindStringSubmatch(userToCheck.Email)
	if m == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Wrong email syntax"))
		return
	}
	if len(userToCheck.Password) < 8 {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("Password to short"))
		return
	}
	q := `select * from user where userEmail = ?`
	row := db.QueryRow(q, userToCheck.Email)
	err = row.Err()
	if err == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Header().Set("Conten-Type", "application/text")
		w.Write([]byte("Sorry, your email already exist"))
		return
	} else if err != sql.ErrNoRows {
		log.Println("Error during checking Email against existing ones")
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Conten-Type", "application/text")
		w.Write([]byte("Sorry, a problem during checking your email"))
		return
	}

	q = `insert into user(userEmail,userPassword)
	values
	(?,?);
	`
	res, err := db.Exec(q, &userToCheck.Email, &userToCheck.Password)
	if err != nil {
		log.Println("Error during inserting to db, ", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Conten-Type", "application/text")
		w.Write([]byte("Sorry, a problem during adding your info to our db"))
	}
	log.Println(res.RowsAffected())
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Conten-Type", "application/text")
	w.Write([]byte("Sucess!"))

}
