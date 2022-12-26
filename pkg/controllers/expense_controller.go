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

	return context.JSON(http.StatusCreated, result)
}

func (c ExpenseController) GetExpenseController(context echo.Context) error {
	id := context.Param("id")

	result, _ := c.ExpenseService.GetExpense(id)

	return context.JSON(http.StatusOK, result)
}

func (c ExpenseController) GetExpensesController(context echo.Context) error {
	result, _ := c.ExpenseService.GetExpenses()

	return context.JSON(http.StatusOK, result)
}

type RequestExpense struct {
	Title  string   `json:"title"`
	Amount float64  `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

func (c ExpenseController) UpdateExpenseController(context echo.Context) error {
	id := context.Param("id")
	requestExpense := RequestExpense{}
	err := context.Bind(&requestExpense)
	if err != nil {
		return context.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	e := entities.Expense{
		Title:  requestExpense.Title,
		Amount: requestExpense.Amount,
		Note:   requestExpense.Note,
		Tags:   requestExpense.Tags,
	}

	result, _ := c.ExpenseService.UpdateExpense(id, e)

	return context.JSON(http.StatusOK, result)
}
