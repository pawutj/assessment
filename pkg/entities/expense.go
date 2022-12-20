package entities

type Expense struct {
	Title  string
	Amount float64
	Note   string
	Tags   []string
}
