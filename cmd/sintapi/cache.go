package main

import "time"

type cache struct {
	products  map[int][]Product
	validTill time.Time
}

func (c *cache) Add(id int, products []Product) {
	c.products[id] = products
	c.validTill = time.Now().Add(time.Minute * 15)
}

func (c *cache) Retrieve(id int) ([]Product, bool) {
	val, ok := c.products[id]
	return val, ok
}
