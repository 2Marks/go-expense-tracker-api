package categories

import (
	"net/http"

	"github.com/2marks/go-expense-tracker-api/middlewares"
	"github.com/2marks/go-expense-tracker-api/types"
	"github.com/2marks/go-expense-tracker-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	service types.CategoryService
}

func NewHandler(service types.CategoryService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.create).Methods(http.MethodPost)
	router.HandleFunc("", h.getAll).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", h.update).Methods(http.MethodPut)
	router.HandleFunc("/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusBadRequest, err)
		return
	}

	var payload = new(types.CreateCategoryDTO)
	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.service.Create(userId, *payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusCreated,
		"category created successfully",
		nil,
	)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusForbidden, err)
		return
	}

	payload := types.GetAllCategoryDTO{
		Page:    utils.GetRequestQueryIntVal(r, "page", 1),
		PerPage: utils.GetRequestQueryIntVal(r, "perPage", 10),
	}

	categories, err := h.service.GetAll(userId, payload)
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"categories fetched successfully",
		*categories,
	)

}
func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusBadRequest, err)
		return
	}

	var payload = new(types.UpdateCategoryDTO)
	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.service.Update(userId, *payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"category updated successfully",
		nil,
	)
}
func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusBadRequest, err)
		return
	}

	payload := types.DeleteCategoryDTO{
		ID: utils.GetReqPathIntVal(r, "id"),
	}

	if err := h.service.Delete(userId, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"category deleted successfully",
		nil,
	)
}
