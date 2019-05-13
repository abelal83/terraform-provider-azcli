package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

// Client comment
type Client struct {
	CosmosAccountName string
}

type account struct {
	State string `json:"state"`
	Name  string `json:"name"`
}

// NewClient returns a new PowerDNS client
func NewClient(cosmosAccountName string) *Client {

	args := []string{"account", "show"}

	cmd := exec.Command("az", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("az account: %q\n", out.String())

	jsonStr := `
	{
	  "data": {
		"object": "card",
		"id": "card_123",
		"last4": "4242"
	  }
	}
	`
	//s, err := strconv.Unquote(out.String())
	s, err := strconv.Unquote(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	var a account
	err = json.Unmarshal([]byte(s), &a)
	if err != nil {

	}

	client := Client{
		CosmosAccountName: cosmosAccountName,
	}
	return &client
}

// AZCommand Run az commands
func (*Client) AZCommand(c []string) (o []byte) {

	out, err := exec.Command("az", c...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("returned data %s\n", out)

	return out

}
