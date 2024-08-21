package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("db> ")
		text, _ := reader.ReadString('\n')
		command := strings.TrimSpace(text)

		if strings.HasPrefix(command, ".") {
			err := doMetaCommand(command)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			}

			continue
		}

		statement, err := prepareStatement(command)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}

		execute_statement(statement)
	}
}

const (
	STATEMENT_INSERT string = "statement_insert"
	STATEMENT_SELECT string = "statement_select"
)

type Statement struct {
	Type string
}

func doMetaCommand(command string) error {
	if command == ".exit" {
		os.Exit(0)
	}
	return fmt.Errorf("unrecognized commmand: %s", command)
}

func prepareStatement(command string) (Statement, error) {
	switch command {
	case "insert":
		return Statement{Type: STATEMENT_INSERT}, nil
	case "select":
		return Statement{Type: STATEMENT_SELECT}, nil
	default:
		return Statement{}, fmt.Errorf("unrecognized statment: %s", command)
	}
}

func execute_statement(statement Statement) {
	switch statement.Type {
	case STATEMENT_INSERT:
		fmt.Println("This would be an insert.")
	case STATEMENT_SELECT:
		fmt.Println("This would be a select.")
	}
}
