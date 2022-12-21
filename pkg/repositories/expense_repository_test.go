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
	want := entities.Expense{Title: "SomeTitle", Amount: 20.0, Note: "SomeNote"}

	result, err := suite.repository.GetExpense(give)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), result.Title, want.Title)
	assert.Equal(suite.T(), result.Amount, want.Amount)
	assert.Equal(suite.T(), result.Note, want.Note)

}

func TestExpenseRepositorySuite(t *testing.T) {
	suite.Run(t, new(ExpenseRepositorySuite))
}
