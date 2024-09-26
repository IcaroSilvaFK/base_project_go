package main

import (
	"log"
	"os"

	"github.com/labstrack/echo/v4"
  "github.com/labstrack/echo/v4/middleware"
  "github.com/joho/godotenv"
)

func main() {

  if err := godotenv.Load(); err != nil {
    log.Fatal(err.Error())
  }

  e := echo.New()
  
  e.Use(middleware.CORS())
  e.Use(middleware.Recover())
  e.Use(middleware.Logger())
  //TODO import routes in app

  port := os.Getenv("PORT")

  if port == "" {
    port = ":8080"
  }

  e.Logger.Fatal(e.Start(port))
}
