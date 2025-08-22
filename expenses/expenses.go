package expenses

import (
	"fmt"
	"ExpenseTracker/models"
	"ExpenseTracker/storage"
)

var expensesFile = "expenses.json"

func AddExpense(username string) {
	expenses := storage.LoadJSON[models.Expense](expensesFile)
	categories := []string{"Travel", "Food", "Fun", "Work"}
	fmt.Println("=== Add Expense ===")
	fmt.Println("Choose category:")
	for i, c := range categories {
		fmt.Printf("%d. %s\n", i+1, c)
	}

	var choice int
	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)
	if choice < 1 || choice > len(categories) {
		fmt.Println("Invalid category.")
		return
	}

	var amount float64
	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)

	expenses = append(expenses, models.Expense{
		Username: username,
		Category: categories[choice-1],
		Amount:   amount,
	})
	storage.SaveJSON(expensesFile, expenses)
	fmt.Println("Expense added successfully!")
}

func ViewReport(username string) {
	expenses := storage.LoadJSON[models.Expense](expensesFile)
	fmt.Println("=== Expense Report ===")
	total := 0.0
	categoryTotals := map[string]float64{}

	for _, e := range expenses {
		if e.Username == username {
			fmt.Printf("%s - %.2f\n", e.Category, e.Amount)
			total += e.Amount
			categoryTotals[e.Category] += e.Amount
		}
	}

	if total == 0 {
		fmt.Println("No expenses found.")
		return
	}

	fmt.Printf("Total spent: %.2f\n", total)

	maxCategory := ""
	maxAmount := 0.0
	for cat, amt := range categoryTotals {
		if amt > maxAmount {
			maxAmount = amt
			maxCategory = cat
		}
	}
	fmt.Printf("You spend the most on: %s (%.2f)\n", maxCategory, maxAmount)
}
