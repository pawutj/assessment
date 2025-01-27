package repositories

import (
	"database/sql"

	"github.com/lib/pq"

	"github.com/pawutj/assessment/pkg/entities"
)

type IExpenseRepository interface {
	CreateExpense(expense entities.Expense) (entities.Expense, error)
	UpdateExpense(id string, expense entities.Expense) (entities.Expense, error)
	GetExpense(id string) (entities.Expense, error)
	GetExpenses() ([]entities.Expense, error)
}

type ExpenseRepository struct {
	DB *sql.DB
}

func (r ExpenseRepository) CreateExpense(expense entities.Expense) (entities.Expense, error) {

	row := r.DB.QueryRow("INSERT INTO EXPENSE (title, amount,note,tags) values ($1, $2 , $3, $4)  RETURNING id,title, amount,note,tags", expense.Title, expense.Amount, expense.Note, pq.Array(expense.Tags))

	var id string
	var title string
	var amount float64
	var note string
	var tags []string
	err := row.Scan(&id, &title, &amount, &note, pq.Array(&tags))

	if err != nil {
		return entities.Expense{}, err
	}

	return entities.Expense{ID: id, Title: title, Amount: amount, Note: note, Tags: tags}, nil
}

func (r ExpenseRepository) GetExpense(id string) (entities.Expense, error) {
	stmt, err := r.DB.Prepare("SELECT id, title, amount , note , tags FROM EXPENSE where id=$1")

	if err != nil {
		return entities.Expense{}, err
	}

	rowId := id
	row := stmt.QueryRow(rowId)

	var _id string
	var amount float64
	var title, note string
	var tags []string

	err = row.Scan(&_id, &title, &amount, &note, pq.Array(&tags))

	if err != nil {
		return entities.Expense{}, err
	}

	return entities.Expense{ID: _id, Title: title, Amount: amount, Note: note, Tags: tags}, nil

}

func (r ExpenseRepository) GetExpenses() ([]entities.Expense, error) {
	stmt, err := r.DB.Prepare("SELECT  id, title, amount , note , tags FROM EXPENSE")
	if err != nil {
		return []entities.Expense{}, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return []entities.Expense{}, err
	}

	expenses := []entities.Expense{}

	for rows.Next() {
		e := entities.Expense{}
		err := rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			return []entities.Expense{}, err
		}
		expenses = append(expenses, e)
	}

	return expenses, nil
}

func (r ExpenseRepository) UpdateExpense(id string, e entities.Expense) (entities.Expense, error) {
	stmt, err := r.DB.Prepare("UPDATE EXPENSE SET title = $2, amount = $3 , note = $4, tags = $5 WHERE id = $1  RETURNING id,title, amount,note,tags")
	if err != nil {
		return entities.Expense{}, err
	}

	row := stmt.QueryRow(id, e.Title, e.Amount, e.Note, pq.Array((e.Tags)))

	var result entities.Expense
	err = row.Scan(&result.ID, &result.Title, &result.Amount, &result.Note, pq.Array(&result.Tags))

	if err != nil {
		return entities.Expense{}, err
	}

	return result, nil
}
