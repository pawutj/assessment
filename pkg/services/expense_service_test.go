package services_test

import (
	"testing"

	"github.com/pawutj/assessment/pkg/entities"
	"github.com/pawutj/assessment/pkg/services"
)

type StubExpensesRepository struct {
	Expense  entities.Expense
	Expenses []entities.Expense
}

func (s StubExpensesRepository) CreateExpense(e entities.Expense) (entities.Expense, error) {
	return s.Expense, nil
}

func (s StubExpensesRepository) GetExpense(id int) (entities.Expense, error) {
	return s.Expense, nil
}

func (s StubExpensesRepository) GetExpenses() ([]entities.Expense, error) {
	return s.Expenses, nil
}

func (s StubExpensesRepository) UpdateExpense(id int, e entities.Expense) (entities.Expense, error) {
	return s.Expense, nil
}

func TestCreateShouldReturnExpense(t *testing.T) {
	give := entities.Expense{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}}
	want := entities.Expense{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}}

	repository := StubExpensesRepository{want, []entities.Expense{}}
	ExpenseService := services.ExpenseService{repository}

	result, err := ExpenseService.CreateExpense(give)

	if err != nil {
		t.Errorf("Error should be nil")
	}

	if result.Title != want.Title {
		t.Errorf("Want '%s' got '%s'", result.Title, want.Title)
	}

}

func TestGetShouldReturnExpense(t *testing.T) {

	give := 0
	want := entities.Expense{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}}

	repository := StubExpensesRepository{want, []entities.Expense{}}
	ExpenseService := services.ExpenseService{repository}

	result, err := ExpenseService.GetExpense(give)

	if err != nil {
		t.Errorf("Error should be nil")
	}

	if result.Title != want.Title {
		t.Errorf("Want '%s' got '%s'", result.Title, want.Title)
	}

}

func TestGetShouldReturnExpenses(t *testing.T) {

	want := []entities.Expense{
		{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}},
		{Title: "Some Deposit", Amount: 10, Note: "Some Note", Tags: []string{"tag1", "tag2"}},
	}

	repository := StubExpensesRepository{entities.Expense{}, want}
	ExpenseService := services.ExpenseService{repository}

	result, err := ExpenseService.GetExpenses()

	if err != nil {
		t.Errorf("Error should be nil")
	}

	if len(result) != len(want) {
		t.Errorf("Want Length '%d' got '%d'", len(want), len(result))
	}

}
