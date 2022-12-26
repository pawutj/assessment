package services_test

import (
	"testing"

	"github.com/pawutj/assessment/pkg/entities"
	"github.com/pawutj/assessment/pkg/services"
	"github.com/stretchr/testify/assert"
)

type StubExpensesRepository struct {
	Expense  entities.Expense
	Expenses []entities.Expense
}

func (s StubExpensesRepository) CreateExpense(e entities.Expense) (entities.Expense, error) {
	return s.Expense, nil
}

func (s StubExpensesRepository) GetExpense(id string) (entities.Expense, error) {
	return s.Expense, nil
}

func (s StubExpensesRepository) GetExpenses() ([]entities.Expense, error) {
	return s.Expenses, nil
}

func (s StubExpensesRepository) UpdateExpense(id string, e entities.Expense) (entities.Expense, error) {
	return s.Expense, nil
}

func TestCreateShouldReturnExpense(t *testing.T) {
	give := entities.Expense{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}}
	want := entities.Expense{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}}

	repository := StubExpensesRepository{want, []entities.Expense{}}
	ExpenseService := services.ExpenseService{repository}

	result, err := ExpenseService.CreateExpense(give)

	assert.Nil(t, err)
	assert.EqualValues(t, want, result)

}

func TestGetShouldReturnExpense(t *testing.T) {

	give := "0"
	want := entities.Expense{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}}

	repository := StubExpensesRepository{want, []entities.Expense{}}
	ExpenseService := services.ExpenseService{repository}

	result, err := ExpenseService.GetExpense(give)

	assert.Nil(t, err)
	assert.EqualValues(t, want, result)

}

func TestGetShouldReturnExpenses(t *testing.T) {

	want := []entities.Expense{
		{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}},
		{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}},
	}

	repository := StubExpensesRepository{entities.Expense{}, want}
	ExpenseService := services.ExpenseService{repository}

	result, err := ExpenseService.GetExpenses()

	assert.Nil(t, err)
	assert.EqualValues(t, want, result)

}

func TestUpdateShouldReturnExpense(t *testing.T) {
	give := entities.Expense{
		Title: "Update", Amount: 10, Note: "Update", Tags: []string{"tag1"},
	}
	id := "2"
	want := entities.Expense{
		Title: "Update", Amount: 10, Note: "Update", Tags: []string{"tag1"},
	}

	repository := StubExpensesRepository{give, []entities.Expense{}}
	ExpenseService := services.ExpenseService{repository}

	result, err := ExpenseService.UpdateExpense(id, give)
	assert.Nil(t, err)
	assert.Equal(t, want, result)
}
