package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawutj/assessment/pkg/entities"
)

func GetExpenseController(c echo.Context) error {
	//id := c.Param("id")
	result := entities.Expense{Title: "Title"}
	return c.JSON(http.StatusOK, result)
}
