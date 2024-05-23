package auth

import (
	"net/http"

	"github.com/2marks/go-expense-tracker-api/types"
	"github.com/2marks/go-expense-tracker-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	service types.AuthService
}

func NewHandler(service types.AuthService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/auth/signup", h.Signup).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", h.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/forgot-password", h.ForgotPassword).Methods(http.MethodPost)
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var payload *types.SignupDTO = new(types.SignupDTO)

	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusNoContent, err)
		return
	}

	if err := h.service.Signup(*payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusCreated,
		"user account created successfully",
		nil,
	)

}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var payload *types.LoginDTO = new(types.LoginDTO)

	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusNoContent, err)
		return
	}

	data, err := h.service.Login(*payload)
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"user logged in successfully",
		data,
	)

}

func (h *Handler) ForgotPassword(w http.ResponseWriter, r *http.Request) {

}
