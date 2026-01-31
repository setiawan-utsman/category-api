package main

import (
	"backend-api/database"
	"backend-api/handlers"
	"backend-api/repositories"
	"backend-api/services"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		fmt.Println("Gagal Koneksi Database", err)
		return
	}
	defer db.Close()

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	// Setup routes
	http.HandleFunc("/api/products", productHandler.CategoryHandler)
	http.HandleFunc("/api/products/", productHandler.CategoryIdHandler)

	http.HandleFunc("/api/categories", categoryHandler.CategoryHandler)
	http.HandleFunc("/api/categories/", categoryHandler.CategoryIdHandler)

	// Start server
	fmt.Println("Server Running di localhost:" + config.Port)
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Gagal Running Server", err)
	}
}
