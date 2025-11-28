package main

import (
	"log"
	"github.com/joho/godotenv"
	"Messanger/cmd/handlers"
	"Messanger/internal/database"
)

func main(){
	if err := godotenv.Load(); err != nil{ // Проверка на подключение к файлу среды окружения .env
		log.Printf("Can't connect to .env file: %v", err)
	}
	database.InitDB() // Проверка на базу данных
	handlers.HandleRequests() // Поддержка запросов
}
