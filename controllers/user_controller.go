package controllers

import (
	"encoding/json"
	"net/http"

	"eCommerceAPI/services"
)

type UserController struct {
	Service *services.UserService
}

func (c *UserController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	type NewUserInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var u NewUserInfo

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	newUser, err := c.Service.RegisterUser(u.Email, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
