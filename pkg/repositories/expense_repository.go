package repositories

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/pawutj/assessment/pkg/entities"
)

type IExpenseRepository interface {
	CreateExpense(expenses entities.Expense) (entities.Expense, error)
	// UpdateExpenses(expenses entities.Expense) (entities.Expense, error)
	GetExpense(id string) (entities.Expense, error)
	// DeleteExpenses(id string) (entities.Expense, error)
}

type ExpenseRepository struct {
	DB *sql.DB
}

func (r ExpenseRepository) CreateExpense(expenses entities.Expense) (entities.Expense, error) {
	return entities.Expense{}, nil
}

func (r ExpenseRepository) GetExpense(id string) (entities.Expense, error) {
	stmt, err := r.DB.Prepare("SELECT id, title, amount , note FROM EXPENSE where id=$1")

	if err != nil {
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := id
	row := stmt.QueryRow(rowId)

	var _id, amount float64
	var title, note string

	err = row.Scan(&_id, &title, &amount, &note)
	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	return entities.Expense{Title: title, Amount: amount, Note: note}, nil

}
