package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/soniccyclone/go-learning/controllers"
	"github.com/soniccyclone/go-learning/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

type Config struct {
	Port             string
	ConnectionString string
}

func NewApp(connectionString string) (App, error) {
	var a App
	err := a.setupDb(connectionString)
	if err == nil {
		a.setupRoutes()
	}
	return a, err
}

// Connects to a DB and migrates entities.
func (a *App) setupDb(connectionString string) error {
	var err error = nil
	a.DB, err = gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err == nil {
		log.Println("Connected to Database...")
		a.DB.AutoMigrate(&entities.Product{})
		log.Println("Database Migration Completed...")
	}
	return err
}

func (a *App) setupRoutes() {
	// Initialize the router
	a.Router = mux.NewRouter().StrictSlash(true)
	c := &controllers.ProductController{a.DB}
	// Register Routes
	a.Router.HandleFunc("/api/products", c.GetProducts).Methods("GET")
	a.Router.HandleFunc("/api/products/{id}", c.GetProductById).Methods("GET")
	a.Router.HandleFunc("/api/products", c.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/api/products/{id}", c.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/api/products/{id}", c.DeleteProduct).Methods("DELETE")
}

func (a *App) Start(port string) {
	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), a.Router))
}
