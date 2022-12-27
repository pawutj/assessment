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

type DummyError struct {
}

func (d DummyError) Error() string {
	return "DummyError"
}

type DummyExpensesErrorRepository struct {
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

func (d DummyExpensesErrorRepository) CreateExpense(e entities.Expense) (entities.Expense, error) {
	return entities.Expense{}, DummyError{}
}

func (d DummyExpensesErrorRepository) UpdateExpense(id string, e entities.Expense) (entities.Expense, error) {
	return entities.Expense{}, DummyError{}
}

func (d DummyExpensesErrorRepository) GetExpense(id string) (entities.Expense, error) {
	return entities.Expense{}, DummyError{}
}

func (d DummyExpensesErrorRepository) GetExpenses() ([]entities.Expense, error) {
	return []entities.Expense{}, DummyError{}

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

func TestCreateError(t *testing.T) {
	repository := DummyExpensesErrorRepository{}
	ExpenseService := services.ExpenseService{repository}

	_, err := ExpenseService.CreateExpense(entities.Expense{})

	assert.NotNil(t, err)
}

func TestGetError(t *testing.T) {
	repository := DummyExpensesErrorRepository{}
	ExpenseService := services.ExpenseService{repository}

	_, err := ExpenseService.GetExpense("0")

	assert.NotNil(t, err)
}

func TestGetAllError(t *testing.T) {
	repository := DummyExpensesErrorRepository{}
	ExpenseService := services.ExpenseService{repository}

	_, err := ExpenseService.GetExpenses()

	assert.NotNil(t, err)

}

func TestUpdateError(t *testing.T) {
	repository := DummyExpensesErrorRepository{}
	ExpenseService := services.ExpenseService{repository}

	_, err := ExpenseService.UpdateExpense("0", entities.Expense{})

	assert.NotNil(t, err)

}
