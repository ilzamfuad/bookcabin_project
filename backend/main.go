package main

import (
	"bookcabin_project/config"
	"bookcabin_project/controller"
	"bookcabin_project/repository"
	"bookcabin_project/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	router := gin.Default()

	db := config.InitSQLiteDB()
	defer func() {
		if sqlDB, err := db.DB(); err != nil {
			panic(err)
		} else {
			_ = sqlDB.Close()
		}
	}()

	voucherRepo := repository.NewVoucherRepository(db)
	voucherService := service.NewVoucherService(voucherRepo)
	voucherController := controller.NewVoucherController(voucherService)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	router.POST("/api/check", voucherController.Check)
	router.POST("/api/generate", voucherController.Generate)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(router)

	server := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Println("Server is running on port", server)
	http.ListenAndServe(server, handler)

}
