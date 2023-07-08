package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"teknight_bc/db"

	"github.com/gorilla/mux"
)

type users struct {
	unic_id string
	name    string
	toc     string
}

type user struct {
	name string
}

// func doesuserExists(w http.ResponseWriter, r *http.Request) {
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var user User
// 	json.Unmarshal(reqBody, &user)
// 	json.NewEncoder(w).Encode(user)
// }

func DoesUserExists(w http.ResponseWriter, r *http.Request) {
	var doesExist bool
	vars := mux.Vars(r)
	name := vars["name"]
	conn := db.DbConn()
	query := "select name from user where name= ?"
	var tag user
	err := conn.QueryRow(query, name).Scan(&tag.name)
	if err != nil {
		doesExist = false
	} else {
		doesExist = true
	}
	json.NewEncoder(w).Encode(doesExist)
}

func Users(w http.ResponseWriter, r *http.Request) {
	conn := db.DbConn()
	query := "Select * from User"
	result := db.DbQuery(conn, query)
	var x []string
	for result.Next() {
		var tag users
		// for each row, scan the result into our tag composite object
		err := result.Scan(&tag.unic_id, &tag.name, &tag.toc)
		if err != nil {
			fmt.Print("no result")
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		x = append(x, tag.name)
		// log.Printf("name:" + tag.name + " with id:" + tag.unic_id + " created at:" + tag.toc)
	}
	result.Close()
	// jsonBody, _ := json.Marshal(x)
	// return jsonBody
	json.NewEncoder(w).Encode(x)
}

func Root(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("ROOT")
}
