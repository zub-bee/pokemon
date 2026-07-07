package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/zub-bee/pokemon/repl"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}

func commandExit() error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	if err != nil {
		os.Exit(1)
		return err
	}
	os.Exit(0)
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		command := repl.CleanInput(userInput)[0]
		switch command {
		case "exit":
			err := commands["exit"].callback()
			if err != nil {
				fmt.Println(err)
			}
		case "help":
			fmt.Println(`
Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
		default:
			fmt.Println("Unknown command")
		}

	}
}
