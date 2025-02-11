package main

import (
	"fmt"
	"log"
	"os"

	"inventory-management/controllers"
	"inventory-management/models"
	"inventory-management/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	appPort := os.Getenv("APP_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.Produk{}, &models.Inventaris{}, &models.Pesanan{})

	r := gin.Default()

	produkController := &controllers.ProdukController{DB: db}
	inventarisController := &controllers.InventarisController{DB: db}
	pesananController := &controllers.PesananController{DB: db}

	routes.SetupRoutes(r, produkController, inventarisController, pesananController)

	r.Run(fmt.Sprintf(":%s", appPort))
}
