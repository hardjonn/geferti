package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnError(t *testing.T) {
	_, err := New(".", ".gibberish-name", "dotenv")

	assert.NotEqual(t, err, nil, "should return error")
}
