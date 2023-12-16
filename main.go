package main 

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)
type User struct {
	FirstName string `json: "firstName"`
	LastName string `json: "lastname"`
	Email string `json: "email"`
}
type Profile struct {
	Department string `json:"department"`
	Designation string `json:"designation"`
	Employee User `json: "employee"`
}

function main (){
        router:= mux.NewRouter()

		http.ListenandServe(":5000", router)
		
}