package main

import (
	"fmt"
	"inventory-management-api/internal/app"
	"inventory-management-api/internal/config"
	"inventory-management-api/internal/models"
	"inventory-management-api/internal/routers"
	"log"
	"os"
	"os/exec"
)

func runSetupScript() error {
	cmd := exec.Command("bash", "./scripts/setup.sh")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	fmt.Println("Ensuring database is ready...")
	if err := runSetupScript(); err != nil {
		fmt.Printf("Setup script failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Connecting to DB...")
	db := config.InitDB()

	fmt.Println("Running migrations...")
	if err := db.AutoMigrate(
		&models.ContactInfo{},
		&models.Document{},
		&models.DocumentLine{},
		&models.Partner{},
		&models.WarehouseLocation{},
		&models.User{},
		&models.Product{},
		&models.Warehouse{},
		&models.StockBalance{},
	); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	fmt.Println("Migrations applied.")

	container := app.NewContainer(db)

	r := routers.SetupRouter(container)

	log.Fatal(r.Run(":8080"))
}
