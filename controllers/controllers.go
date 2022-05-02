package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gabriel-Rabeloo/go-rest-api/database"
	"github.com/Gabriel-Rabeloo/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func FindAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var personalities []models.Personality
	database.DB.Order("id asc").Find(&personalities)
	json.NewEncoder(w).Encode(personalities)
}

func FindPersonalityById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Create(&personality)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.Delete(&personality, id)
	w.WriteHeader(http.StatusNoContent)
}

func UpdatePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality

	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Save(&personality)
	w.WriteHeader(http.StatusNoContent)
}
