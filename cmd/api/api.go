package api

import (
	"fmt"
	"net/http"

	"github.com/2marks/go-expense-tracker-api/internal/auth"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type ApiServer struct {
	db   *gorm.DB
	addr string
}

func NewApiServer(db *gorm.DB, addr string) *ApiServer {
	return &ApiServer{db: db, addr: addr}
}

func (a *ApiServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	/** start auth routes */
	authRepository := auth.NewRepository(a.db)
	authService := auth.NewService(authRepository)
	authHandler := auth.NewHandler(authService)
	authHandler.RegisterRoutes(subRouter)
	/** end auth routes */

	fmt.Println("server listening on", a.addr)

	return http.ListenAndServe(a.addr, router)
}
