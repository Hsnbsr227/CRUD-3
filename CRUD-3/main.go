package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ConnDb() 

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Add User")
		fmt.Println("2. List Users")
		fmt.Println("3. Delete User")
		fmt.Println("4. Update User")
		fmt.Println("0. Exit")
		fmt.Print("Your choice: ")

		var secim string
		fmt.Scan(&secim)

		switch secim {
		case "1":

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter name: ")
			
			nameRaw, _ := reader.ReadString('\n')
			name := strings.TrimSpace(nameRaw)
			fmt.Scan(&name)

			fmt.Print("Enter password: ")
			
			passwordRaw, _ := reader.ReadString('\n')
			password := strings.TrimSpace(passwordRaw)
			fmt.Scan(&password)	
			
			err := CreateUser(name, password)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "2":
			ReadUser()

		case "3":
			fmt.Print("Enter the ID of the user to delete: ")
			var id int
			fmt.Scan(&id)

			err := DeleteUser(id)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "4":
			reader := bufio.NewReader(os.Stdin)
			var id int

			fmt.Print("Enter the ID of the user to update: ")
			fmt.Scan(&id)

			fmt.Print("Enter new username: ")
			nameRaw, _ := reader.ReadString('\n')
			newName := strings.TrimSpace(nameRaw)
			fmt.Scan(&newName)

			fmt.Print("Enter new password: ")
			passRaw, _ := reader.ReadString('\n')
			newPassword := strings.TrimSpace(passRaw)
			fmt.Scan(&newPassword)

			err := UpdateUser(id, newName, newPassword)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case "0":
			fmt.Println("The program is being terminated...")
			return

		default:
			fmt.Println("Invalid selection. Please try again.")
		}
	}

}
