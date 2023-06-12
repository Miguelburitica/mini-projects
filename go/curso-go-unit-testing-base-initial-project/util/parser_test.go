package util

import (
	"catching-pokemons/models"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParserPokemonSuccess(t *testing.T) {
	c := require.New(t)

	body, err := ioutil.ReadFile("./samples/pokeapi_response.json")

	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &response)
	c.NoError(err)

	parsedPokemon, err := ParsePokemon(response)
	c.NoError(err)

	body, err = ioutil.ReadFile("./samples/api_response.json")
	c.NoError(err)
	
	var expectedPokemon models.Pokemon

	err = json.Unmarshal(body, &expectedPokemon)
	c.NoError(err)

	c.Equal(expectedPokemon, parsedPokemon)
}

func TestParserPokemonTypeNotFound(t *testing.T) {
	c := require.New(t)

	body, err := ioutil.ReadFile("./samples/pokeapi_response.json")

	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &response)
	c.NoError(err)

	response.PokemonType = []models.PokemonType{}

	_, err = ParsePokemon(response)
	c.NotNil(err)
	c.EqualError(ErrNotFoundPokemonType, err.Error())
}

func BenchmarkParser(b *testing.B) {
	c := require.New(b)

	body, err := ioutil.ReadFile("./samples/pokeapi_response.json")
	c.NoError(err)

	var response models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &response)
	c.NoError(err)

	for i := 0; i < b.N; i++ {
		_, err = ParsePokemon(response)
		c.NoError(err)
	}
}
