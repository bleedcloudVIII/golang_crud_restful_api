package main

import (
	"server/database"
	"server/router"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
	\brief function main - start app
*/

func main() {
	config :=
		database.Config{
			User:     "root",
			Password: "root",
			DB:       "gotest",
		}

	connectionString := database.GetConnectionString(config)
	database.Connect(connectionString)
	database.Migrate()

	router.StartRouter()

}
