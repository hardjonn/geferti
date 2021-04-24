package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnError(t *testing.T) {
	_, err := New(".", ".gibberish-name", "dotenv")

	assert.NotEqual(t, err, nil, "should return error")
}

func TestShouldMakeConfig(t *testing.T) {
	c, err := New(".", "../../.env", "dotenv")

	assert.NoError(t, err)
	assert.NotEmpty(t, c)

	v := reflect.ValueOf(*c)

	for i := 0; i < v.NumField(); i++ {
		value := v.Field(i).Interface()
		assert.NotEmpty(t, value)
	}
}
