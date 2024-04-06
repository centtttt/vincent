package main

import (
	"database/sql"
	"fmt"
	"louisweb/models"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "LouisG_DB:louis12345@tcp(database-1.cbkwoowkgsdi.us-east-1.rds.amazonaws.com:3306)/LouisDB")

	if err != nil {
		fmt.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	fmt.Println("Succesful Connection to Database!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT userName, userAge, userEmail FROM Users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user models.Users
			err := rows.Scan(&user.UserName, &user.UserAge, &user.UserEmail)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "User: %+v\n", user)
		}
	})

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
