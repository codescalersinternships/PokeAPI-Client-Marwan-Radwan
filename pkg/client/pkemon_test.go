package client

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPokeByName(t *testing.T) {
	tests := []struct {
		name           string
		pokemonName    string
		mockResponse   string
		mockStatusCode int
		contentType    string
		expectedError  bool
		expectedResult Pokemon
	}{
		{
			name:           "Success",
			pokemonName:    "pikachu",
			mockResponse:   `{"id": 25, "name": "pikachu", "base_experience": 112, "height": 4}`,
			mockStatusCode: http.StatusOK,
			contentType:    "application/json",
			expectedError:  false,
			expectedResult: Pokemon{ID: 25, Name: "pikachu", BaseExperience: 112, Height: 4},
		},
		{
			name:           "Not Found",
			pokemonName:    "unknown",
			mockResponse:   `Not Found`,
			mockStatusCode: http.StatusNotFound,
			contentType:    "text/plain",
			expectedError:  true,
			expectedResult: Pokemon{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				if _, err := w.Write([]byte(tt.mockResponse)); err != nil {
					t.Errorf("failed to write mock response: %v", err)
				}
			}))
			defer server.Close()

			client := &Client{
				config: Config{URL: server.URL},
				httpClient: &http.Client{
					Transport: &http.Transport{},
				},
			}

			ctx := context.Background()
			result, err := client.GetPokeByName(ctx, tt.pokemonName)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
func TestGetAllPokemon(t *testing.T) {
	tests := []struct {
		name           string
		mockResponse   string
		mockStatusCode int
		contentType    string
		expectedError  bool
		expectedResult PokemonList
	}{
		{
			name:           "Success",
			mockResponse:   `{"results": [{"id": 25, "name": "pikachu", "base_experience": 112, "height": 4}]}`,
			mockStatusCode: http.StatusOK,
			contentType:    "application/json",
			expectedError:  false,
			expectedResult: PokemonList{Results: []Pokemon{{ID: 25, Name: "pikachu", BaseExperience: 112, Height: 4}}},
		},
		{
			name:           "Not Found",
			mockResponse:   `Not Found`,
			mockStatusCode: http.StatusNotFound,
			contentType:    "text/plain",
			expectedError:  true,
			expectedResult: PokemonList{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockStatusCode)
				if _, err := w.Write([]byte(tt.mockResponse)); err != nil {
					t.Errorf("failed to write mock response: %v", err)
				}
			}))
			defer server.Close()

			client := &Client{
				config: Config{URL: server.URL},
				httpClient: &http.Client{
					Transport: &http.Transport{},
				},
			}

			ctx := context.Background()
			result, err := client.GetAllPokemon(ctx)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedResult, result)
			}
		})
	}
}
