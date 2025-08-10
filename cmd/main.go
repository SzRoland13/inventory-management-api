package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	dsn := "host=localhost user=postgres password=postgres dbname=your_db_name port=5432 sslmode=disable TimeZone=Europe/Budapest"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // debug info a konzolra
	})

	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	router := gin.Default()

	fmt.Println("Backend started on :8080")
	router.Run(":8080")
}
