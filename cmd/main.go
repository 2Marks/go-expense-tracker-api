package main

import (
	"fmt"
	"log"

	"github.com/2marks/go-expense-tracker-api/cmd/api"
	"github.com/2marks/go-expense-tracker-api/config"
	"github.com/2marks/go-expense-tracker-api/database"
)

func main() {
	newDatabase := database.NewDatabase()
	db := newDatabase.InitDatabase()
	newDatabase.RunMigrations(db)
	newDatabase.RunSeeds(db)

	apiServer := api.NewApiServer(db, fmt.Sprintf(":%s", config.Envs.Port))
	if err := apiServer.Run(); err != nil {
		log.Fatalf("error starting up server: %s", err.Error())
	}
}
