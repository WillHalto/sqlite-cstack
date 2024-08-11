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
		if command == "exit" {
			return
		} else {
			fmt.Printf("Unrecognized command: %s\n", command)
		}
	}
}
