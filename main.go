package main

import (
	"log"
	"os"
	"time"

	"github.com/iqbal13/tugas3hactiv/config"
	"github.com/iqbal13/tugas3hactiv/models"
	"github.com/iqbal13/tugas3hactiv/routes"
	"github.com/joho/godotenv"
)

func main() {

	os.Setenv("TZ", "Asia/Jakarta")
	time.LoadLocation("Asia/Jakarta")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed Load Environment")
	}
	config.ConnectDB()
	config.DB.AutoMigrate(&models.WaterWind{})
	app := routes.SetupRouter()
	app.Run(os.Getenv("APP_PORT"))

}
