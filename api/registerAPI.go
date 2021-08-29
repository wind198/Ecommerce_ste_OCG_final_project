package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	//todo
	//get user data from request body
	//create a userdata struct and decode request body into it
	//connect to database and check if a user with such email already exist
	//if not exist add it to db and anwser to response
	db := newDbHanlder.db
	var userToRegister userToCheck
	err := json.NewDecoder(r.Body).Decode(&userToRegister)
	log.Println(userToRegister)
	if err != nil {
		log.Println("Error during decoding request body", err)
		return
	}
	var validEmail = regexp.MustCompile(`(\w+\d*)@(\w+)\.(net|com|org)`)
	m := validEmail.FindStringSubmatch(userToRegister.UserEmail)
	if m == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("wrong email syntax"))
		return
	}
	if len(userToRegister.UserPassword) < 8 {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte("password to short"))
		return
	}
	q := `select userEmail from user where userEmail = ?`
	row := db.QueryRow(q, userToRegister.UserEmail)
	var userAlreadyExist userToCheck
	err = row.Scan(&userAlreadyExist.UserEmail)
	if err == nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Header().Set("Conten-Type", "application/text")
		w.Write([]byte("email already exist"))
		return
	} else if err != sql.ErrNoRows {
		log.Println("Error during checking Email against existing ones")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	q = `insert into user(userEmail,userPassword)
	values
	(?,?);
	`
	res, err := db.Exec(q, &userToRegister.UserEmail, &userToRegister.UserPassword)
	if err != nil {
		log.Println("Error during inserting to db, ", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	log.Println(res.RowsAffected())
	w.WriteHeader(http.StatusCreated)
}
