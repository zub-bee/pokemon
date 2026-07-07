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
	"mapb": {
		name:        "mapb",
		description: "display pokemon map backwards",
		callback:    showMapB,
	},
}

var conf = &config{
	Previous: "",
	Next:     "",
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
	_, err := fmt.Println("displaying pokemon map")
	if err != nil {
		return err
	}

	url := "https://pokeapi.co/api/v2/location-area"

	if conf.Next != "" {
		url = conf.Next
	}

	// get response from api
	response, err := getResponse(url)
	if err != nil {
		return err
	}

	// resolve the response to config and lists
	// config
	conf.Previous = response.Previous
	conf.Next = response.Next

	// lists
	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func showMapB() error {
	_, err := fmt.Println("displaying pokemon map")
	if err != nil {
		return err
	}

	url := "https://pokeapi.co/api/v2/location-area"

	if conf.Previous != "" {
		url = conf.Previous
	}

	// get response from api
	response, err := getResponse(url)
	if err != nil {
		return err
	}

	// resolve the response to config and lists
	// config
	conf.Previous = response.Previous
	conf.Next = response.Next

	// lists
	for _, area := range response.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func getResponse(url string) (response, error) {
	res, err := http.Get(url)
	if err != nil {
		return response{}, err
	}
	defer res.Body.Close()

	var resp response

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resp)
	if err != nil {
		return response{}, err
	}

	return resp, nil
}

func printHelp() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, command := range commands {
		fmt.Println(command.name + ": " + command.description)
	}
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
			printHelp()
		case "map":
			commands["map"].callback()
		case "mapb":
			commands["mapb"].callback()
		default:
			fmt.Println("Unknown command")
		}

	}
}
