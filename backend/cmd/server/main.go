package main

import (
	"banking-app/backend/internal/bank"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Initialize repository with data directory
	dataDir := "../../db"
	repo := bank.NewRepository(dataDir)

	// Initialize service and handler
	service := bank.NewService(repo)
	handler := bank.NewHandler(service)

	fmt.Println("Welcome to Bank Management CLI!")
	fmt.Println("Type 'help' for available commands or 'exit' to quit.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("bankster > ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		// Parse command
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "exit", "quit":
			fmt.Println("Goodbye!")
			return

		case "help":
			handler.ShowHelp()

		case "create":
			if len(parts) < 2 {
				fmt.Println("Usage: create <name>")
				continue
			}
			handler.HandleCreate(parts[1])

		case "get":
			if len(parts) < 2 {
				fmt.Println("Usage: get <id>")
				continue
			}
			handler.HandleGet(parts[1])

		case "list":
			handler.HandleList()

		case "update":
			if len(parts) < 3 {
				fmt.Println("Usage: update <id> <name>")
				continue
			}
			handler.HandleUpdate(parts[1], parts[2])

		case "delete":
			if len(parts) < 2 {
				fmt.Println("Usage: delete <id>")
				continue
			}
			handler.HandleDelete(parts[1])

		case "add-customer":
			if len(parts) < 3 {
				fmt.Println("Usage: add-customer <bank_id> <customer_id>")
				continue
			}
			handler.HandleAddCustomer(parts[1], parts[2])

		case "remove-customer":
			if len(parts) < 3 {
				fmt.Println("Usage: remove-customer <bank_id> <customer_id>")
				continue
			}
			handler.HandleRemoveCustomer(parts[1], parts[2])

		case "get-customers":
			if len(parts) < 2 {
				fmt.Println("Usage: get-customers <bank_id>")
				continue
			}
			handler.HandleGetCustomers(parts[1])

		default:
			fmt.Printf("Unknown command: %s. Type 'help' for available commands.\n", command)
		}

		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
}
