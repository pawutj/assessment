package service

import (
	"github.com/pawutj/assessment/pkg/entities"
	"github.com/pawutj/assessment/pkg/repositories"
)

type ExpenseService struct {
	ExpensesRepository repositories.IExpensesRepository
}

func (s ExpenseService) CreateExpenses(e entities.Expenses) (entities.Expenses, error) {
	result, err := s.ExpensesRepository.CreateExpenses(e)
	if err != nil {
		return entities.Expenses{}, err
	}
	return result, nil
}

func (s ExpenseService) GetExpenses(id string) (entities.Expenses, error) {
	result, err := s.ExpensesRepository.GetExpenses(id)
	if err != nil {
		return entities.Expenses{}, err
	}
	return result, nil
}
