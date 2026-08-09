package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//go:embed index.html
var indexHTML []byte

type AbilityDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Ability struct {
	Ability  AbilityDetail `json:"ability"`
	IsHidden bool          `json:"is_hidden"`
	Slot     int           `json:"slot"`
}

type TypeDetail struct {
	Name string `json:"name"`
}

type TypeSlot struct {
	Slot int        `json:"slot"`
	Type TypeDetail `json:"type"`
}

type StatDetail struct {
	Name string `json:"name"`
}

type Stat struct {
	BaseStat int        `json:"base_stat"`
	Stat     StatDetail `json:"stat"`
}

type OfficialArtwork struct {
	FrontDefault string `json:"front_default"`
}

type OtherSprites struct {
	OfficialArtwork OfficialArtwork `json:"official-artwork"`
}

type Sprites struct {
	FrontDefault string       `json:"front_default"`
	Other        OtherSprites `json:"other"`
}

type Cries struct {
	Latest string `json:"latest"`
	Legacy string `json:"legacy"`
}

type Pokemon struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	BaseExperience int        `json:"base_experience"`
	Height         int        `json:"height"`
	Weight         int        `json:"weight"`
	Sprites        Sprites    `json:"sprites"`
	Types          []TypeSlot `json:"types"`
	Abilities      []Ability  `json:"abilities"`
	Stats          []Stat     `json:"stats"`
	Cries          Cries      `json:"cries"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/index.html" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(indexHTML)
}

func pokemonAPIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	queryName := r.URL.Query().Get("name")
	pokemonName := strings.ToLower(strings.TrimSpace(queryName))
	if pokemonName == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Pokémon name is required"})
		return
	}

	apiURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", url.PathEscape(pokemonName))
	resp, err := http.Get(apiURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to fetch data from PokeAPI"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("Pokémon %q not found. Please check spelling!", pokemonName)})
		return
	}

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(resp.StatusCode)
		json.NewEncoder(w).Encode(ErrorResponse{Error: fmt.Sprintf("PokeAPI returned status %d", resp.StatusCode)})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to read API response"})
		return
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Failed to parse API response"})
		return
	}

	json.NewEncoder(w).Encode(pokemon)
}

func main() {
	http.HandleFunc("/api/pokemon", pokemonAPIHandler)
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("⚡ PokéApp server running on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
