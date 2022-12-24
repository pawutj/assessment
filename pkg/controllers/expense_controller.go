package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pawutj/assessment/pkg/entities"
	"github.com/pawutj/assessment/pkg/services"
)

type Err struct {
	Message string `json:"message"`
}

type ExpenseController struct {
	ExpenseService services.IExpenseService
}

func (c ExpenseController) CreateExpenseController(context echo.Context) error {
	e := entities.Expense{}
	err := context.Bind(&e)

	if err != nil {
		return context.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	result, err := c.ExpenseService.CreateExpense(e)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	// result := entities.Expense{Title: "Title"}
	return context.JSON(http.StatusOK, result)
}

func (c ExpenseController) GetExpenseController(context echo.Context) error {
	id := context.Param("id")
	result, _ := c.ExpenseService.GetExpense(id)
	// result := entities.Expense{Title: "Title"}
	return context.JSON(http.StatusOK, result)
}

func (c ExpenseController) GetExpensesController(context echo.Context) error {
	result, _ := c.ExpenseService.GetExpenses()
	// result := entities.Expense{Title: "Title"}
	return context.JSON(http.StatusOK, result)
}
