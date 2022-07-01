package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_add(t *testing.T) {
	sum := add(5, 5)
	assert.Equal(t, sum, 10)
}
