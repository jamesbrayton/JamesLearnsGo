package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	example1 := 1
	example2 := 2

	assert.NotEqual(t, example1, example2)
}
