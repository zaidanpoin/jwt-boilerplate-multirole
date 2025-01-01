package main

import (
	"fmt"
	"log"

	"github.com/zaidanpoin/blog-go/Database"
	"github.com/zaidanpoin/blog-go/Model"
	"github.com/zaidanpoin/blog-go/Router"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	fmt.Println("asad")

	loadDatabase()
	Router.ServeApps()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {

	Database.Connect()
	Database.Database.AutoMigrate(&Model.User{})
}
