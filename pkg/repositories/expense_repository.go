package repositories

import (
	"github.com/pawutj/assessment/pkg/entities"
)

type IExpenseRepository interface {
	CreateExpense(expenses entities.Expense) (entities.Expense, error)
	// UpdateExpenses(expenses entities.Expense) (entities.Expense, error)
	GetExpense(id string) (entities.Expense, error)
	// DeleteExpenses(id string) (entities.Expense, error)
}
