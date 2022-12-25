package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/pawutj/assessment/pkg/controllers"
	"github.com/pawutj/assessment/pkg/db"
	"github.com/pawutj/assessment/pkg/entities"
	"github.com/pawutj/assessment/pkg/repositories"
	"github.com/pawutj/assessment/pkg/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExpenseIntegralSuite struct {
	suite.Suite
	e echo.Echo
}

func (suite *ExpenseIntegralSuite) SetupSuite() {
	expenseController := controllers.ExpenseController{
		ExpenseService: services.ExpenseService{
			ExpenseRepository: repositories.ExpenseRepository{DB: db.ConnectDB()},
		},
	}

	suite.e = *echo.New()

	go func(e *echo.Echo) {
		suite.e.GET("/expenses", expenseController.GetExpensesController)
		suite.e.GET("/expenses/:id", expenseController.GetExpenseController)
		suite.e.POST("/expenses", expenseController.CreateExpenseController)
		suite.e.PUT("/expenses/:id", expenseController.UpdateExpenseController)
		suite.e.Start(":" + os.Getenv("PORT"))
	}(&suite.e)

}

func (suite *ExpenseIntegralSuite) TearDownSuite() {
}

func TestExpenseRepositorySuite(t *testing.T) {
	suite.Run(t, new(ExpenseIntegralSuite))
}

func (suite *ExpenseIntegralSuite) TestGetExpenseByID() {

	var e entities.Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(1)), nil)
	// res := request(http.MethodGet, uri("expenses", "1"), nil)
	err := res.Decode(&e)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
	assert.Equal(suite.T(), e.Title, "SomeTitle")

}

func (suite *ExpenseIntegralSuite) TestGetAllExpense() {

	var es []entities.Expense
	res := request(http.MethodGet, uri("expenses"), nil)
	// res := request(http.MethodGet, uri("expenses", "1"), nil)
	err := res.Decode(&es)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
	assert.Greater(suite.T(), len(es), 1)

}

func (suite *ExpenseIntegralSuite) TestPostExpense() {

	body := bytes.NewBufferString(`{
		"title": "SomeTitle",
		"amount": 20.0,
		"Note":"SomeNote",
		"tags": ["tag1"]
	}`)

	var e entities.Expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&e)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
	assert.Equal(suite.T(), e.Title, "SomeTitle")

}

func (suite *ExpenseIntegralSuite) TestUpdateExpense() {

	body := bytes.NewBufferString(`{
		"title": "SomeTitleUpdate",
		"amount": 20.0,
		"Note":"SomeNote",
		"tags": ["tag1"]
	}`)

	var e entities.Expense

	res := request(http.MethodPut, uri("expenses", strconv.Itoa(2)), body)
	err := res.Decode(&e)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, res.StatusCode)
	assert.Equal(suite.T(), e.Title, "SomeTitleUpdate")
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	return json.NewDecoder(r.Body).Decode(v)
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	// req.Header.Add("Authorization", os.Getenv("AUTH_TOKEN"))
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}
