package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawutj/assessment/pkg/services"
)

type ExpenseController struct {
	ExpenseService services.IExpenseService
}

func (c ExpenseController) GetExpenseController(context echo.Context) error {
	id := context.Param("id")
	result, _ := c.ExpenseService.GetExpense(id)
	// result := entities.Expense{Title: "Title"}
	return context.JSON(http.StatusOK, result)
}
