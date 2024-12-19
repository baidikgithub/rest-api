package main

import (
	"log"
	"rest-api/handler"
	"rest-api/models"
	"rest-api/repository"
	"rest-api/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "user=postgres password=Shohom@789 dbname=db_employee sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Employee{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	repo := repository.NewEmployeeRepository(db)
	serv := service.NewEmployeeService(repo)
	h := handler.NewEmployeeHandler(serv)

	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/users", h.GetAllEmployees)
	e.GET("/users/:id", h.GetEmployeeByID)
	e.POST("/users", h.AddEmployee)
	e.PUT("/users/update/:id", h.UpdateEmployee)
	e.DELETE("/users/delete/:id", h.DeleteEmployee)

	log.Println("Server is running on port 8088")
	e.Logger.Fatal(e.Start(":8088"))
}
