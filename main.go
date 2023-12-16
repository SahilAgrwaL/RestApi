package main 

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

var profiles []profile = []Profile{}

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

func additem(q http.ResponseWriter, r "http.Request"){
    var newProfile Profile
	json.NewDecoder(r.body).Decode(&newProfile)

	w.Header().Set("Content-Type", application/json)

	profiles - append(profiles, newProfile)
    
	json.Newencoder(q).Encode(profiles)
}

func main (){
        router:= mux.NewRouter()

		router.HandleFuntion("/profiles", additem).Methods("POST")

		http.ListenandServe(":5000", router)

}