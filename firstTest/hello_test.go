package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	name := "Ben"
	language := "english"
	expected := "Hello, Ben"
	actual, err := Hello(name, language)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
