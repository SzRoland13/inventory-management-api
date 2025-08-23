package main

import (
	"fmt"
	"inventory-management-api/internal/app"
	"inventory-management-api/internal/config"
	"inventory-management-api/internal/routers"
	"log"
)

func main() {
	fmt.Println("Connecting to DB...")
	db := config.InitDB()

	container := app.NewContainer(db)

	r := routers.SetupRouter(container)

	log.Fatal(r.Run(":8080"))
}
