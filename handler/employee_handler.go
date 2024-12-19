package handler

import (
	"net/http"
	"rest-api/models"
	"rest-api/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EmployeeHandler struct {
	service service.EmployeeService
}

func NewEmployeeHandler(service service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) GetAllEmployees(c echo.Context) error {
	employees, err := h.service.GetAllEmployees()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch employees"})
	}
	return c.JSON(http.StatusOK, employees)
}

func (h *EmployeeHandler) GetEmployeeByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := h.service.GetEmployeeByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Employee not found"})
	}
	return c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) AddEmployee(c echo.Context) error {
	var employee models.Employee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	err := h.service.AddEmployee(&employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add employee"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Employee added successfully"})
}

func (h *EmployeeHandler) UpdateEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var employee models.Employee
	if err := c.Bind(&employee); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}
	employee.ID = id
	err := h.service.UpdateEmployee(&employee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update employee"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Employee updated successfully"})
}

func (h *EmployeeHandler) DeleteEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.DeleteEmployee(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete employee"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Employee deleted successfully"})
}
