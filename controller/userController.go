package controller

import (
	"bookingApp/model"
	"encoding/json"
	"net/http"
	"sync"
)

var cacheMutex sync.RWMutex

func FindAllUsers (
	w http.ResponseWriter,
	r *http.Request,
) {
	cacheMutex.RLock()
	users := model.FindAll()
	cacheMutex.RUnlock()

	response := map[string]any{
		"users": users,
		"links": []map[string]string{
			{
				"rel":    "self",
				"href":   "/users",
				"method": "GET",
			},
			{
				"rel":    "create",
				"href":   "/users",
				"method": "POST",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func CreateUser (
	w http.ResponseWriter,
	r *http.Request,
) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheMutex.Lock()
	id, err := model.Create(user)
	if err != nil {
		switch err {
		case model.ErrUserExists:
			http.Error(w, err.Error(), http.StatusConflict)
			cacheMutex.Unlock()
			return
		case model.ErrInvalidUser:
			http.Error(w, err.Error(), http.StatusBadRequest)
			cacheMutex.Unlock()
			return
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			cacheMutex.Unlock()
			return
		}	
	}
	cacheMutex.Unlock()

	response := map[string]interface{}{
		"id":   id,
		"links": []map[string]string{
			{
				"rel":    "self",
				"href":   "/users/" + id.String(),
				"method": "GET",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", "/user/"+id.String())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}