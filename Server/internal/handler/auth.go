package handler

import (
	"Server/internal/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/mail"
	"regexp"
	"unicode/utf8"
)

var (
	hasUpper   = regexp.MustCompile(`[A-Z]`)
	hasLower   = regexp.MustCompile(`[a-z]`)
	hasDigit   = regexp.MustCompile(`[0-9]`)
	hasSpecial = regexp.MustCompile(`[#?!@$%^&*-]`)
	isASCII    = regexp.MustCompile(`^[\x00-\x7F]+$`)
)

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{auth: auth}
}

type signUpRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type signInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req signUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if !req.validate() {
		SendError(w, http.StatusBadRequest, "not all fields valid")
		return
	}

	token, err := h.auth.SignUp(r.Context(), req.Email, req.FirstName, req.LastName, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrEmailTaken) {
			SendError(w, http.StatusConflict, service.ErrEmailTaken.Error())
			return
		}

		SendError(w, http.StatusInternalServerError, "internal server error")
		log.Println(err)
		return
	}

	SendSuccess(w, http.StatusCreated, authResponse{Token: token})
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var req signInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendError(w, http.StatusBadRequest, "invalid request")
		return
	}

	if !req.validate() {
		SendError(w, http.StatusBadRequest, "not all fields valid")
		return
	}

	token, err := h.auth.SignIn(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			SendError(w, http.StatusUnauthorized, service.ErrInvalidCredentials.Error())
			return
		}

		SendError(w, http.StatusInternalServerError, "internal server error")
		log.Println(err)
		return
	}

	SendSuccess(w, http.StatusCreated, authResponse{Token: token})
}

func (r *signUpRequest) validate() bool {
	// Check if email is valid
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return false
	}

	// Check if name is valid
	if utf8.RuneCountInString(r.FirstName) > 40 || utf8.RuneCountInString(r.FirstName) <= 0 {
		return false
	}
	if utf8.RuneCountInString(r.LastName) > 40 || utf8.RuneCountInString(r.LastName) <= 0 {
		return false
	}

	// Check password
	if !isValidPassword(r.Password) {
		return false
	}

	return true
}

func (r *signInRequest) validate() bool {
	// Check if email is valid
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return false
	}

	return true
}

func isValidPassword(password string) bool {
	return len(password) >= 8 &&
		len(password) <= 32 &&
		isASCII.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasLower.MatchString(password) &&
		hasDigit.MatchString(password) &&
		hasSpecial.MatchString(password)
}
