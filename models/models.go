package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Expense struct {
	Username string  `json:"username"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
}
