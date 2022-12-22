package services

import (
	"github.com/pawutj/assessment/pkg/entities"
	"github.com/pawutj/assessment/pkg/repositories"
)

type IExpenseService interface {
	CreateExpense(e entities.Expense) (entities.Expense, error)
	GetExpense(id string) (entities.Expense, error)
}

type ExpenseService struct {
	ExpenseRepository repositories.IExpenseRepository
}

func (s ExpenseService) CreateExpense(e entities.Expense) (entities.Expense, error) {
	result, err := s.ExpenseRepository.CreateExpense(e)

	if err != nil {
		return entities.Expense{}, err
	}
	return result, nil
}

func (s ExpenseService) GetExpense(id string) (entities.Expense, error) {
	result, err := s.ExpenseRepository.GetExpense(id)
	if err != nil {
		return entities.Expense{}, err
	}
	return result, nil
}
