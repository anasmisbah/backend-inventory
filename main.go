package main

import (
	"log"
	"os"

	"github.com/anasmisbah/backend-inventory/config"
	"github.com/anasmisbah/backend-inventory/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	config.InitDB()
	e := echo.New()
	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(getPort()))
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":"+envPort
	}
	return "3000";
}

func loadEnv()  {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}