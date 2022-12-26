package repositories

import (
	"testing"

	"github.com/pawutj/assessment/pkg/db"
	"github.com/pawutj/assessment/pkg/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExpenseRepositorySuite struct {
	suite.Suite
	repository ExpenseRepository
}

func (suite *ExpenseRepositorySuite) SetupSuite() {
	DB := db.ConnectDB()
	suite.repository = ExpenseRepository{DB}

}

func (suite *ExpenseRepositorySuite) TestSQL() {
	_, err := suite.repository.DB.Query("SELECT * FROM Expense WHERE TITLE = 'SomeTitle'")
	assert.Nil(suite.T(), err)
}

func (suite *ExpenseRepositorySuite) TestGetExpense() {

	give := "1"

	result, err := suite.repository.GetExpense(give)

	assert.Nil(suite.T(), err)

	assert.NotEqual(suite.T(), result.Title, "")
	assert.NotEqual(suite.T(), result.Amount, 0)
	assert.NotEqual(suite.T(), result.Note, "")

}

func (suite *ExpenseRepositorySuite) TestCreateExpense() {

	give := entities.Expense{Title: "SomeTitle", Amount: 20.0, Note: "SomeNote", Tags: []string{"tags1"}}
	want := entities.Expense{Title: "SomeTitle", Amount: 20.0, Note: "SomeNote", Tags: []string{"tags1"}}

	result, err := suite.repository.CreateExpense(give)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), result.Title, want.Title)
	assert.Equal(suite.T(), result.Amount, want.Amount)
	assert.Equal(suite.T(), result.Note, want.Note)
	assert.Equal(suite.T(), result.Tags, want.Tags)

}

func (suite *ExpenseRepositorySuite) TestUpdateExpense() {
	id := "2"
	give := entities.Expense{Title: "SomeTitle1", Amount: 20.0, Note: "SomeNote", Tags: []string{"tags1"}}
	want := entities.Expense{Title: "SomeTitle1", Amount: 20.0, Note: "SomeNote", Tags: []string{"tags1"}}

	result, err := suite.repository.UpdateExpense(id, give)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), result.Title, want.Title)
	assert.Equal(suite.T(), result.Amount, want.Amount)
	assert.Equal(suite.T(), result.Note, want.Note)
	assert.Equal(suite.T(), result.Tags, want.Tags)

}

func (suite *ExpenseRepositorySuite) TestGetExpenses() {

	result, err := suite.repository.GetExpenses()

	assert.Nil(suite.T(), err)
	assert.Greater(suite.T(), len(result), 1)

}

func TestExpenseRepositorySuite(t *testing.T) {
	suite.Run(t, new(ExpenseRepositorySuite))
}
