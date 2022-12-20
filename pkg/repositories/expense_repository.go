package repositories

import (
	"github.com/pawutj/assessment/pkg/entities"
)

type IExpensesRepository interface {
	CreateExpenses(expenses entities.Expenses) (entities.Expenses, error)
	UpdateExpenses(expenses entities.Expenses) (entities.Expenses, error)
	GetExpenses(id string) (entities.Expenses, error)
	DeleteExpenses(id string) (entities.Expenses, error)
}
