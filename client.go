package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/tidwall/gjson"
)

// Client comment
type Client struct {
	CosmosAccountName string
}

type account struct {
	State string `json:"state"`
	Name  string `json:"name"`
}

// NewClient AZ Client
func NewClient(cosmosAccountName string) *Client {

	args := []string{"account", "show"}
	out, err := exec.Command("az", args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	output := string(out)

	if gjson.Valid(output) {
		log.Print("az cli output is valid json")
	} else {
		panic("az cli output not valid json")
	}

	value := gjson.Get(output, "name")
	fmt.Print(value)

	client := Client{
		CosmosAccountName: cosmosAccountName,
	}
	return &client
}

// AZCommand Run az commands
func (c Client) AZCommand(cmd []string) string {

	out, err := exec.Command("az", cmd...).Output()
	if err != nil {
		log.Fatal(err)
	}

	output := string(out)
	return output

}
