package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetPokemonFromPokeApiSucces(t *testing.T) {
	c := require.New(t)

	pokemon, err := GetPokemonFromPokeApi("blastoise")
	c.NoError(err)

	body, err := ioutil.ReadFile("../util/samples/pokeapi_response.json")
	c.NoError(err)
	
	var expectedPokemon models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &expectedPokemon)
	c.NoError(err)

	c.Equal(expectedPokemon, pokemon)
}

func TestGetPokemonFromPokeApiSuccessWithMocks(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "blastoise"
	
	body, err := ioutil.ReadFile("../util/samples/pokeapi_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(200, string(body)))

	pokemon, err := GetPokemonFromPokeApi(id)
	c.NoError(err)

	var expectedPokemon models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &expectedPokemon)
	c.NoError(err)

	c.Equal(expectedPokemon, pokemon)
}

func TestGetPokemonInternalError(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "blastoise"
	
	body, err := ioutil.ReadFile("../util/samples/pokeapi_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(500, string(body)))

	_, err = GetPokemonFromPokeApi(id)
	c.NotNil(err)
	c.EqualError(ErrPokeApiFailed, err.Error())
}

func TestGetPokemon(t *testing.T) {
	c := require.New(t)

	r, err := http.NewRequest(http.MethodGet, "/pokemon/{id}", nil)	
	c.NoError(err)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "blastoise",
	}

	r = mux.SetURLVars(r, vars)

	GetPokemon(w, r)

	expectedBodyResponse, err := ioutil.ReadFile("../util/samples/api_response.json")
	c.NoError(err)

	var expectedResponse models.Pokemon

	err = json.Unmarshal(expectedBodyResponse, &expectedResponse)
	c.NoError(err)

	var actualPokemon models.Pokemon

	err = json.Unmarshal(w.Body.Bytes(), &actualPokemon)
	c.NoError(err)

	c.Equal(http.StatusOK, w.Code)
	c.Equal(expectedResponse, actualPokemon)
}
