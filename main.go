package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"vincentweb/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "vincentdb:vincentdb26@tcp(database-1.cgfuegyd9ftf.us-east-1.rds.amazonaws.com)/vincentdb")
	
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

	fmt.Println("Succesfully connect to Database!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT agentName, agentType, agentHP, agentAbility FROM agentDetail")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var agent models.AgentDetails
			err := rows.Scan(&agent.AgentName, &agent.AgentType, &agent.AgentHP, &agent.AgentAbility)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Fprintf(w, "Agent: %+v\n", agent)
		}
	})

	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
