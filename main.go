package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"eCommerceAPI/controllers"
	"eCommerceAPI/middleware"
	"eCommerceAPI/repository"
	"eCommerceAPI/services"

	// The underscore means we import it to register the driver,
	// but we won't call its functions directly in our code.
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: no .env file found")
	}

	// Read the variables
	connStr := os.Getenv("DB_CONN_STR")
	if connStr == "" {
		log.Fatal("DB_CONN_STR enviroment variable is required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Fallback if empty
	}

	// Database connection to setup
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error configuring database: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}
	log.Println("Successfully connected to PostgreSQL!")

	// Initialize the layers(Dependency Injection)
	productRepo := &repository.ProductRepo{DB: db}
	productService := &services.ProductService{Repo: productRepo}
	productController := &controllers.ProductController{Service: productService}

	orderRepo := &repository.OrderRepo{DB: db}
	orderService := &services.OrderService{Repo: orderRepo}
	orderController := &controllers.OrderController{Service: orderService}

	authController := &controllers.AuthController{}

	// Public Routes(No auth required)
	http.HandleFunc("POST /login", middleware.Logging(authController.LoginHandler))
	http.HandleFunc("GET /products", middleware.Logging(productController.GetProductsHandler))

	// Protected Routes
	http.HandleFunc("POST /products", middleware.Logging(middleware.RequireAuth(productController.CreateProductHandler)))
	http.HandleFunc("POST /checkout", middleware.Logging(middleware.RequireAuth(orderController.CheckoutHandler)))
	http.HandleFunc("GET /orders", middleware.Logging(middleware.RequireAuth(orderController.GetOrdersHandler)))

	// Start the server
	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server error: ", err)
	}
}
