# Pokedex CLI

An interactive command-line Pokedex written in Go. The application uses
[PokéAPI](https://pokeapi.co/) to browse location areas, discover Pokémon, and
build a local collection.

This project demonstrates Go fundamentals including HTTP requests, JSON
decoding, generic functions, command-line input, package organization,
concurrency-safe state, and automated tests.

## Features

- Browse Pokémon location areas with pagination
- Explore a location to see the Pokémon found there
- Attempt to catch Pokémon with randomized outcomes
- Inspect the stats and types of caught Pokémon
- List the Pokémon in your local Pokedex
- Normalize and parse interactive command input
- Cache API responses with an expiring, mutex-protected in-memory cache

## Requirements

- Go 1.26.4 or newer
- Internet access for requests to PokéAPI

## Getting started

Clone the repository and run the CLI:

```bash
git clone https://github.com/zub-bee/pokemon.git
cd pokemon
go run .
```

You can also build a standalone executable:

```bash
go build -o pokemon .
./pokemon
```

## Commands

| Command | Description |
| --- | --- |
| `help` | Show available commands |
| `map` | Display the next page of location areas |
| `mapb` | Display the previous page of location areas |
| `explore <location>` | List Pokémon found in a location area |
| `catch <pokemon>` | Attempt to catch a Pokémon |
| `inspect <pokemon>` | Show details for a caught Pokémon |
| `pokedex` | List caught Pokémon |
| `exit` | Close the application |

Example session:

```text
Pokedex > map
Pokedex > explore canalave-city-area
Pokedex > catch pikachu
Pokedex > inspect pikachu
Pokedex > pokedex
```

## Testing

Run all tests with:

```bash
go test ./...
```

The test suite covers command input normalization and cache behavior,
including expiration of cached values.

## Project structure

```text
.
├── main.go                  # CLI, API integration, and command handlers
├── repl/                    # Input normalization helpers and tests
├── internal/pokecache/      # Thread-safe expiring cache and tests
└── go.mod                   # Go module definition
```

## API

Data is provided by the public [PokéAPI](https://pokeapi.co/docs/v2). No API
key is required.
