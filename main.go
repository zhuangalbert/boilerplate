package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/zhuangalbert/boilerplate/src/api/databases"
	"github.com/zhuangalbert/boilerplate/src/api/v1/controllers"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	cmdString := command()

	if cmdString == "" {
		cmdString = "serve"
	}

	if cmdString == "serve" && os.Getenv("APP_ENV") == "production" {
		startProd()
		startApp()
	} else if cmdString == "serve" && os.Getenv("APP_ENV") != "production" {
		startApp()
	} else if cmdString == "seed" {
		//database.Seed()
	} else if cmdString == "migrate" {
		databases.Migrate()
	}
}

func startProd() {

}

func startApp() {
	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	router := gin.Default()
	controllers.UserControllerHandler(router)
	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)
	router.Run(serverString)
}

func command() string {
	args := os.Args[1:]
	if len(args) > 0 {
		return args[0]
	}
	return ""
}
