package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_cache_Add(t *testing.T) {
	c := cache{
		products:  map[int][]Product{},
		validTill: time.Time{},
	}

	c.Add(1, []Product{{ID: 1}})
	assert.Len(t, c.products, 1)
}

func Test_cache_Retrieve(t *testing.T) {
	c := cache{
		products:  map[int][]Product{},
		validTill: time.Time{},
	}

	c.Add(1, []Product{{ID: 1}})

	products, ok := c.Retrieve(1)

	assert.True(t, ok)
	assert.Len(t, products, 1)
}
