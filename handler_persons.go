package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ChukwunonsoFrank/hngx-task-two/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (newAPIConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	person, err := newAPIConfig.DB.CreatePerson(r.Context(), database.CreatePersonParams{
		ID: uuid.New(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user resource: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(person))
}

func (newAPIConfig *apiConfig) handlerFetchUser(w http.ResponseWriter, r *http.Request) {
	personIdString := chi.URLParam(r, "user_id")
	person_id, err := uuid.Parse(personIdString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing user ID: %v", err))
		return
	}

	person, err := newAPIConfig.DB.FetchPerson(r.Context(), person_id)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error fetching user details: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(person))
}

func (newAPIConfig *apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	type Parameters struct {
		Name string `json:"name"`
		ID string `json:"id"`
	}

	personIdString := chi.URLParam(r, "user_id")
	person_id, err := uuid.Parse(personIdString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing user ID: %v", err))
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := Parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	person, err := newAPIConfig.DB.UpdatePerson(r.Context(), database.UpdatePersonParams{
		Name: params.Name,
		ID: person_id,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user resource: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(person))
}

func (newAPIConfig *apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	personIdString := chi.URLParam(r, "user_id")
	person_id, err := uuid.Parse(personIdString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing user ID: %v", err))
		return
	}

	person, err := newAPIConfig.DB.DeletePerson(r.Context(), person_id)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error deleting user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(person))
}