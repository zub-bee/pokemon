package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zub-bee/pokemon/repl"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      *config
}

type config struct {
	Previous string
	Next     string
}

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type response struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "display pokemon map",
		callback:    showMap,
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

func showMap() error {
	//_, err := fmt.Println("displaying pokemon map")
	//if err != nil {
	//	return err
	//}

	// get response from api
	url := "https://pokeapi.co/api/v2/location-area"

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var response response

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return err
	}

	// resolve the response to config and lists
	// config
	config := &config{}
	config.Previous = response.Previous
	config.Next = response.Next

	// lists
	for _, area := range response.Results {
		println(area.Name)
	}

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
		case "map":
			commands["map"].callback()
		default:
			fmt.Println("Unknown command")
		}

	}
}
