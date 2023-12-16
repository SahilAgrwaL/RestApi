package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	w.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(w).Encode(profiles)
}

func getAllProfiles(q http.ResponseWriter, r *http.Request){
	q.Header().Set("Content-Type","application/json")
	json.NewEncoder(q).Encode(profiles)
}

func getProfile(q http.ResponseWriter, r *http.Request){
	var idParam string = mux.Vars(r)["id"]
    id, err := strconv.Atoi(idParam)
	if err!=nil{
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id>=len(profiles){
		q.WriteHeader(404)
		q.Write([]byte("No profile found with specified ID"))
		return
	}
	profile := profiles[id]
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profile)
}

func updateProfile(q http.ResponseWriter, r *http.Request){
	var idParam string = mux.Vars(r)["id"]
    id, err := strconv.Atoi(idParam)
	if err!=nil{
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to integer"))
		return
	}
	if id>=len(profiles){
		q.WriteHeader(404)
		q.Write([]byte("No profile found with specified ID"))
		return
	}
	var updateProfile Profile
	json.NewDecoder(r.Body).Decode(&updateProfile)

	profiles[id]=updateProfile
	q.Header().Set("Content-Type","application/json")
	json.NewEncoder(q).Encode(updateProfile)
}

func deleteProfile(q http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		q.WriteHeader(http.StatusBadRequest)
		q.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id >= len(profiles) || id < 0 {
		q.WriteHeader(http.StatusNotFound)
		q.Write([]byte("No profile found with specified ID"))
		return
	}

	// Remove the element at the specified index
	profiles = append(profiles[:id], profiles[id+1:]...)

	q.WriteHeader(http.StatusOK)
}

// func deleteProfile(q http.ResponseWriter, r *http.Request){
// 	var idParam string = mux.Vars(r)["id"]
//     id, err := strconv.Atoi(idParam)
// 	if err!=nil{
// 		q.WriteHeader(400)
// 		q.Write([]byte("ID could not be converted to integer"))
// 		return
// 	}
// 	if id>=len(profiles){
// 		q.WriteHeader(404)
// 		q.Write([]byte("No profile found with specified ID"))
// 		return
// 	}
	
// 	profiles = append(profiles[:id], profiles[:id+1:]... )
// 	q.WriteHeader(200)
// 	}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")
	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")
    router.HandleFunc("/profiles/{id}", getProfile).Methods("GET")
	router.HandleFunc("/profiles/{id}", updateProfile).Methods("PUT")
	router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")
	
	http.ListenAndServe(":5000", router)
}