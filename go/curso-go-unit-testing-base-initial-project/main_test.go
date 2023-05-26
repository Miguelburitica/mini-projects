package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddSuccess(t *testing.T) {
	c := require.New(t)
	
	result := Add(21, 2)

	expect := 23
	c.Equal(expect, result)
}
