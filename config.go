package main

import (
	"log"
)

// Config required for az cli
type Config struct {
	CosmosAccountName string
}

// Client returns a new client for az cli
func (c *Config) Client() (*Client, error) {
	client := NewClient(c.CosmosAccountName)
	log.Printf("[INFO] Cosmos Client configured for server %s", c.CosmosAccountName)
	return client, nil
}
