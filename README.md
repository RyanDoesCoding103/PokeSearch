# ⚡ PokéSearch Explorer (Go + PokéAPI)

A modern, fast, and responsive Pokémon Search web application built with **Go (`net/http`)** backend and vanilla HTML5, CSS3, and JavaScript frontend.

![PokéSearch Explorer](https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/882.png)

## ✨ Features

- **Go REST Server**: Built using Go's standard library `net/http` to query [PokéAPI](https://pokeapi.co/).
- **Pokémon Cries Audio**: Plays official Pokémon cry sound effects (Modern Latest and Classic Legacy audio).
- **Rich Card Visuals**: Official artwork, type badges with dynamic color themes, animated stat meters, base experience, height, and weight.
- **Glassmorphism UI**: High-end dark theme with smooth micro-animations and quick-search pills.
- **Robust Error Handling**: Graceful error management for invalid inputs or 404 responses.

## 🚀 Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.18 or higher installed on your machine.

### Run Locally

1. Clone the repository:
   ```bash
   git clone <YOUR_REPOSITORY_URL>
   cd pokemon-search-app
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

## 📁 Project Structure

```text
Lesson 2/
├── main.go      # Go web server & PokéAPI proxy handler
├── index.html   # Single page app with glassmorphism UI & audio player
└── README.md    # Documentation
```

## 📜 License

MIT License.
