package expenses

import (
	"net/http"

	"github.com/2marks/go-expense-tracker-api/middlewares"
	"github.com/2marks/go-expense-tracker-api/types"
	"github.com/2marks/go-expense-tracker-api/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	expenseService types.ExpenseService
}

func NewHandler(service types.ExpenseService) *Handler {
	return &Handler{expenseService: service}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("", h.create).Methods(http.MethodPost)
	router.HandleFunc("", h.getAll).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", h.getOne).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", h.update).Methods(http.MethodPut)
	router.HandleFunc("/{id:[0-9]+}", h.delete).Methods(http.MethodDelete)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusBadRequest, err)
		return
	}

	var payload *types.CreateExpenseDTO = new(types.CreateExpenseDTO)
	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := h.expenseService.Create(userId, *payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusCreated,
		"expense created successfully",
		nil,
	)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusForbidden, err)
		return
	}

	payload := types.GetAllExpensesDTO{
		Page:    utils.GetRequestQueryIntVal(r, "page", 1),
		PerPage: utils.GetRequestQueryIntVal(r, "perPage", 10),
	}

	expenses, err := h.expenseService.GetAll(userId, payload)
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"expenses fetched successfully",
		*expenses,
	)
}

func (h *Handler) getOne(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusForbidden, err)
		return
	}

	payload := types.GetOneExpenseDTO{
		ID: utils.GetReqPathIntVal(r, "id"),
	}

	expense, err := h.expenseService.GetById(userId, payload)
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"expense fetched successfully",
		expense,
	)
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusForbidden, err)
		return
	}

	var payload = new(types.UpdateExpenseDTO)
	if err := utils.ParseRequestBody(r, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}
	payload.ID = utils.GetReqPathIntVal(r, "id")

	if err := h.expenseService.Update(userId, *payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"expense updated successfully",
		nil,
	)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	userId, err := middlewares.GetUserFromContext(r.Context())
	if err != nil {
		utils.WriteErroresponseToJson(w, http.StatusForbidden, err)
		return
	}

	payload := types.DeleteExpenseDTO{
		ID: utils.GetReqPathIntVal(r, "id"),
	}

	if err := h.expenseService.Delete(userId, payload); err != nil {
		utils.WriteErroresponseToJson(w, http.StatusUnprocessableEntity, err)
		return
	}

	utils.WriteSuccessResponseToJson(
		w,
		http.StatusOK,
		"expense deleted successfully",
		nil,
	)
}
