package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApplication(t *testing.T) {
	app := NewApplication()
	assert.Equal(t, app.sources["bollie"].GetHost().String(), "http://mockserver:3001")
	assert.Equal(t, app.sources["coolbere"].GetHost().String(), "http://mockserver:3002")
	assert.Equal(t, app.sources["aliblabla"].GetHost().String(), "http://mockserver:3003")
}
