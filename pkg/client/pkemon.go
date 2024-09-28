package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	retry "github.com/codescalersinternships/PokeAPI-Client-Marwan-Radwan/internal/Retry"
)

// Pokemon represents the response data from the API
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
}

// PokemonList represents all pokemons
type PokemonList struct {
	Results []Pokemon `json:"results"`
}

// GetPokeByName retrieves a pokemon data using its name
func (c *Client) GetPokeByName(ctx context.Context, pokemonName string) (Pokemon, error) {
	var pokemon Pokemon
	operation := func() error {
		url := fmt.Sprintf("%s/%s", c.config.URL, pokemonName)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}

		req.Header.Set("Accept", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("response failed: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return nil
	}

	if err := retry.Retry(operation); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

// GetAllPokemon retrieves a list of all pokemons
func (c *Client) GetAllPokemon(ctx context.Context) (PokemonList, error) {
	var pokemons PokemonList

	operation := func() error {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.config.URL, nil)
		if err != nil {
			return fmt.Errorf("failed to create request: %w", err)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("response failed: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
		}

		if err := json.NewDecoder(resp.Body).Decode(&pokemons); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return nil
	}

	if err := retry.Retry(operation); err != nil {
		return PokemonList{}, err
	}

	return pokemons, nil
}
