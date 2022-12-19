package repositories

import (
	"github.com/pawutj/assessment/pkg/entities"
)

type IExpensesRepository interface {
	CreateExpenses(expenses entities.Expenses) entities.Expenses
	UpdateExpenses(expenses entities.Expenses) entities.Expenses
	GetExpenses(id string) entities.Expenses
	DeleteExpenses(id string) entities.Expenses
}
