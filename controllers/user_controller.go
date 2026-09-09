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
	_ = json.NewEncoder(w).Encode(newUser)
}

func (c *UserController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	type UserInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	type TokenResponse struct {
		Token string `json:"token"`
	}
	var u UserInfo

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	token, err := c.Service.LoginUser(u.Email, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(TokenResponse{Token: token})
}
