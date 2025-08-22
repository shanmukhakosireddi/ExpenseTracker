package main

import (
	"ExpenseTracker/auth"
	"ExpenseTracker/expenses"
	"ExpenseTracker/models"
	"fmt"
)

func ExpenseMenu(user *models.User) {
	for {
		fmt.Println("\n=== Expense Tracker Menu ===")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Report")
		fmt.Println("3. Logout")
		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			expenses.AddExpense(user.Username)
		case 2:
			expenses.ViewReport(user.Username)
		case 3:
			fmt.Println("Logged out successfully.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func main() {
	for {
		fmt.Println("\n=== Main Menu ===")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Exit")
		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			user := auth.Login()
			if user != nil {
				ExpenseMenu(user)
			}
		case 2:
			auth.Register()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
