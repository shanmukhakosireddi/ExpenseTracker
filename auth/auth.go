package auth

import (
	"fmt"
	"ExpenseTracker/models"
	"ExpenseTracker/storage"
)

var usersFile = "users.json"




// Check if user exists
func UserExists(username string, users []models.User) bool {
	for _, u := range users {
		if u.Username == username {
			return true
		}
	}
	return false
}

// Register a new user
func Register() {
	users := storage.LoadJSON[models.User](usersFile)
	var username, password string
	fmt.Println("=== Register ===")
	fmt.Print("Enter username: ")
	fmt.Scan(&username)

	if UserExists(username, users) {
		fmt.Println("User already exists. Please login.")
		return
	}

	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	users = append(users, models.User{Username: username, Password: password})
	storage.SaveJSON(usersFile, users)
	fmt.Println("Registration successful! Please login now.")
}

// Login a user
func Login() *models.User {
	users := storage.LoadJSON[models.User](usersFile)
	if len(users) == 0 {
		fmt.Println("No users found. Please register first.")
		Register()
		return nil
	}

	var username, password string
	fmt.Println("=== Login ===")
	fmt.Print("Enter username: ")
	fmt.Scan(&username)
	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	for _, u := range users {
		if u.Username == username && u.Password == password {
			fmt.Println("Login successful!")
			return &u
		}
	}
	fmt.Println("Invalid credentials.")
	return nil
}
