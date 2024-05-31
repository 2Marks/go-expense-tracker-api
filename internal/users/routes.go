package users

import (
	"fmt"
	"net/http"

	"github.com/2marks/go-expense-tracker-api/middlewares"
	"github.com/2marks/go-expense-tracker-api/types"
	"github.com/2marks/go-expense-tracker-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	service types.UserService
}

func NewHandler(service types.UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/me", h.GetLoggedInUser).Methods(http.MethodGet)
	router.HandleFunc("/me", h.UpdateUserDetails).Methods(http.MethodPut)
}

func (h *Handler) GetLoggedInUser(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, fmt.Errorf("error while fetching user"))
		return
	}

	user, err := h.service.GetLoggedInUser(userId)
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusNotFound, fmt.Errorf("user not found"))
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"user fetched successfully",
		*user,
	)
}

func (h *Handler) UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, fmt.Errorf("error while fetching user"))
		return
	}

	payload := &types.UpdateUserDetailsDTO{ID: userId}
	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.service.UpdateUserDetails(*payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"user details updated successfully",
		nil,
	)
}
