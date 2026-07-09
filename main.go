package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zub-bee/pokemon/internal/pokecache"
	"github.com/zub-bee/pokemon/repl"
)

type cliCommand struct {
	name        string
	description string
	callback    func(a string) error
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

type MapResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

type ExploreResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"pokemon"`
	VersionDetails []struct {
		Version struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"version"`
		MaxChance        int `json:"max_chance"`
		EncounterDetails []struct {
			MinLevel        int           `json:"min_level"`
			MaxLevel        int           `json:"max_level"`
			ConditionValues []interface{} `json:"condition_values"`
			Chance          int           `json:"chance"`
			Method          struct {
				Name string `json:"name"`
				Url  string `json:"url"`
			} `json:"method"`
		} `json:"encounter_details"`
	} `json:"version_details"`
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
		callback:    showMap,
	},
	"explore": {
		name:        "explore",
		description: "explore pokemon map",
		callback:    exploreMap,
	},
}

var conf = &config{
	Previous: "",
	Next:     "",
}

var cache = &pokecache.Cache{}

func commandExit(a string) error {
	_, err := fmt.Println(a)
	if err != nil {
		os.Exit(1)
		return err
	}
	os.Exit(0)
	return nil
}

func showMap(direction string) error {
	_, err := fmt.Println("displaying pokemon map")
	if err != nil {
		return err
	}

	url := "https://pokeapi.co/api/v2/location-area"

	switch direction {
	case "previous":
		if conf.Previous != "" {
			url = conf.Previous
		}
	default:
		if conf.Next != "" {
			url = conf.Next
		}
	}

	// get response from api
	response, err := GetResponse[MapResponse](url)
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

func exploreMap(name string) error {
	_, err := fmt.Println("exploring pokemon map")
	if err != nil {
		return err
	}

	url := "https://pokeapi.co/api/v2/location-area" + name

	response, err := GetResponse[ExploreResponse](url)
	if err != nil {
		return err
	}

	for _, area := range response.PokemonEncounters {
		fmt.Println(area.Pokemon.Name)
	}

	return nil
}

func GetResponse[T any](url string) (T, error) {
	var resp T

	// check cache first
	fmt.Println("checking cache..")
	data, ok := cache.Get(url)

	if ok {
		fmt.Println("cache hit!")
		return resp, json.Unmarshal(data, &resp)
	}

	fmt.Println("cache miss! calling api")
	// get response from api in cache miss
	res, err := http.Get(url)
	if err != nil {
		return resp, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return resp, err
	}
	fmt.Println("api response received")

	// update cache
	fmt.Println("updating cache..")
	_, ok = cache.Add(url, data)
	if !ok {
		return resp, errors.New("cache could not be updated")
	}

	// unmarshal response
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return resp, err
	}

	// return response
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
			err := commands["exit"].callback("Closing the Pokedex... Goodbye!")
			if err != nil {
				fmt.Println(err)
			}
		case "help":
			printHelp()
		case "map":
			err := commands["map"].callback("next")
			if err != nil {
				fmt.Println(err)
			}
		case "mapb":
			err := commands["mapb"].callback("previous")
			if err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Unknown command")
		}

	}
}
