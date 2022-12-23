package entities

type Expense struct {
	ID     int
	Title  string
	Amount float64
	Note   string
	Tags   []string
}
