package repositories

import (
	"testing"

	"github.com/pawutj/assessment/pkg/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ExpenseRepositorySuite struct {
	suite.Suite
	repository ExpenseRepository
}

func (suite *ExpenseRepositorySuite) SetupSuite() {
	DB := db.ConnectDB()
	suite.repository = ExpenseRepository{DB}

}

func (suite *ExpenseRepositorySuite) TestSQL() {
	_, err := suite.repository.DB.Query("SELECT * FROM Expense WHERE TITLE = 'SomeTitle'")
	assert.Nil(suite.T(), err)

}

func TestExpenseRepositorySuite(t *testing.T) {
	suite.Run(t, new(ExpenseRepositorySuite))
}

// func TestConnectDB(t *testing.T) {
// 	DB := db.ConnectDB()

// 	if DB == nil {
// 		t.Fatalf("ConnectDB failed")
// 	}

// 	_, err := DB.Query("SELECT * FROM Expense WHERE TITLE = 'SomeTitle'")
// 	if err != nil {
// 		t.Errorf("Error %s", err.Error())
// 	}

// 	// for rows.Next() {
// 	// 	var id, price int
// 	// 	var name string

// 	// 	err = rows.Scan(&id, &name, &price)
// 	// 	if err != nil {
// 	// 		t.Fatalf("Error from scan")
// 	// 	}

// 	// 	if name != "SomeProduct" {
// 	// 		t.Errorf("Want SomeProduct result %s", name)
// 	// 	}
// 	// }
// }
