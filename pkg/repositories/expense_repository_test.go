package repositories_test

import (
	"testing"

	"github.com/pawutj/assessment/pkg/db"
)

func TestConnectDB(t *testing.T) {
	DB := db.ConnectDB()

	if DB == nil {
		t.Fatalf("ConnectDB failed")
	}

	_, err := DB.Query("SELECT * FROM Expense WHERE TITLE = 'SomeTitle'")
	if err != nil {
		t.Errorf("Error %s", err.Error())
	}

	// for rows.Next() {
	// 	var id, price int
	// 	var name string

	// 	err = rows.Scan(&id, &name, &price)
	// 	if err != nil {
	// 		t.Fatalf("Error from scan")
	// 	}

	// 	if name != "SomeProduct" {
	// 		t.Errorf("Want SomeProduct result %s", name)
	// 	}
	// }
}
