package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//Get a DB to check user against the database
	db := newDbHanlder.db
	//create a struct to map data from post request
	var userFromRequest userToCheck
	err := json.NewDecoder(r.Body).Decode(&userFromRequest)
	if err != nil {
		log.Println("Error during decoding request body", err)
	}
	//query DB
	q := `select userEmail from user where userEmail = ? and userPassword=?`
	row := db.QueryRow(q, &userFromRequest.UserEmail, &userFromRequest.UserPassword)
	var userOK userToCheck
	err = row.Scan(&userOK.UserEmail)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("Error during query and scan", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response := ResponseToLogin{
		UserEmail: userOK.UserEmail,
	}

	JWTToken := getJWTtoken(userOK.UserEmail)
	response.JWTtoken = JWTToken

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println("Error during mapping json")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJson)
}
func getJWTtoken(userEmail string) string {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		Issuer:    "localhost",
		Audience:  userEmail,
	}
	JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Println(claims)
	JWTSignedString, err := JWTtoken.SignedString([]byte(os.Getenv("HS256_SECRET_KEY")))
	if err != nil {
		log.Panic("Error during signing token, ", err)

	}
	return JWTSignedString
}
