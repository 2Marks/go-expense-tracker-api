package api

import (
	"fmt"
	"net/http"

	"github.com/2marks/go-expense-tracker-api/internal/auth"
	"github.com/2marks/go-expense-tracker-api/internal/categories"
	"github.com/2marks/go-expense-tracker-api/internal/expenses"
	"github.com/2marks/go-expense-tracker-api/internal/users"
	"github.com/2marks/go-expense-tracker-api/middlewares"
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

	/** start auth middleware */
	authMiddleware := middlewares.NewAuthMiddleware()
	/** end auth middleware */

	/** start expense routes */
	expenseRepository := expenses.NewRepository(a.db)
	expenseService := expenses.NewService(expenseRepository)
	expenseHandler := expenses.NewHandler(expenseService)
	expenseRouter := subRouter.PathPrefix("/expenses").Subrouter()
	expenseRouter.Use(authMiddleware.Authenticate)
	expenseHandler.RegisterRoutes(expenseRouter)
	/** end expense routes */

	/* start category routes */
	categoryRepository := categories.NewRepository(a.db)
	categoryService := categories.NewService(categoryRepository)
	categoryHandler := categories.NewHandler(categoryService)
	categoryRouter := subRouter.PathPrefix("/categories").Subrouter()
	categoryRouter.Use(authMiddleware.Authenticate)
	categoryHandler.RegisterRoutes(categoryRouter)
	/** end category routes */

	/** start users routes */
	userRepository := users.NewRepository(a.db)
	userService := users.NewService(userRepository)
	userHandler := users.NewHandler(userService)
	userRouter := subRouter.PathPrefix("/users").Subrouter()
	userRouter.Use(authMiddleware.Authenticate)
	userHandler.RegisterRoutes(userRouter)
	/** end users routes */

	fmt.Println("server listening on", a.addr)

	return http.ListenAndServe(a.addr, router)
}
