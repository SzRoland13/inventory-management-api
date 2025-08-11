package main

import (
	"fmt"
	"inventory-management-api/internal/app"
	"inventory-management-api/internal/config"
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
	fmt.Println("Running setup script...")
	if err := runSetupScript(); err != nil {
		fmt.Printf("Setup script failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Setup finished, starting backend...")

	db := config.InitDB()

	container := app.NewContainer(db)

	r := routers.SetupRouter(container)

	log.Fatal(r.Run(":8080"))
}
