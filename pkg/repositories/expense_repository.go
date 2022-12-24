package repositories

import (
	"database/sql"
	"log"

	"github.com/lib/pq"

	"github.com/pawutj/assessment/pkg/entities"
)

type IExpenseRepository interface {
	CreateExpense(expenses entities.Expense) (entities.Expense, error)
	// UpdateExpenses(expenses entities.Expense) (entities.Expense, error)
	GetExpense(id string) (entities.Expense, error)
	GetExpenses() ([]entities.Expense, error)
	// DeleteExpenses(id string) (entities.Expense, error)
}

type ExpenseRepository struct {
	DB *sql.DB
}

func (r ExpenseRepository) CreateExpense(expense entities.Expense) (entities.Expense, error) {

	row := r.DB.QueryRow("INSERT INTO EXPENSE (title, amount,note,tags) values ($1, $2 , $3, $4)  RETURNING id,title, amount,note,tags", expense.Title, expense.Amount, expense.Note, pq.Array(expense.Tags))

	var id int
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
		log.Fatal("can'tprepare query one row statment", err)
	}

	rowId := id
	row := stmt.QueryRow(rowId)

	var _id int
	var amount float64
	var title, note string
	var tags []string

	err = row.Scan(&_id, &title, &amount, &note, pq.Array(&tags))

	if err != nil {
		log.Fatal("can't Scan row into variables", err)
	}

	return entities.Expense{Title: title, Amount: amount, Note: note, Tags: tags}, nil

}

func (r ExpenseRepository) GetExpenses() ([]entities.Expense, error) {
	stmt, err := r.DB.Prepare("SELECT  id, title, amount , note , tags FROM EXPENSE")
	if err != nil {
		log.Fatal("can'tprepare query one row statment", err.Error())
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all expense: ", err.Error())
	}

	expenses := []entities.Expense{}

	for rows.Next() {
		e := entities.Expense{}
		err := rows.Scan(&e.ID, &e.Title, &e.Amount, &e.Note, pq.Array(&e.Tags))
		if err != nil {
			log.Fatal("can't scan expense: ", err.Error())
		}
		expenses = append(expenses, e)
	}

	return expenses, nil
}
