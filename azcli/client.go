package azcli

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/tidwall/gjson"
)

// Client comment
type Client struct {
	State string `json:"state"`
	Name  string `json:"name"`
	User  user
}

type user struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// NewClient checks az cli client is configured and resturns a new client
func NewClient(s string) *Client {

	args := []string{"account", "set", "--subscription", s}
	out, err := exec.Command("az", args...).Output()
	if err != nil {
		log.Panicf("az cli unable to set subscription %s, error: %s", s, err)
	}

	args = []string{"account", "show"}
	out, err = exec.Command("az", args...).Output()
	if err != nil {
		log.Panicf("az cli not configured %s", err)
	}

	output := string(out)

	if gjson.Valid(output) {
		log.Print("az cli output is valid json")
	} else {
		panic("az cli output not valid json")
	}

	n := gjson.Get(output, "name")
	log.Printf("az cli connected to %s", n)
	var client Client
	json.Unmarshal(out, &client)

	return &client
}

// AZCommand Run az commands
func (c Client) AZCommand(cmd []string) string {

	cmd = append(cmd, "--subscription", c.Name)
	out, err := exec.Command("az", cmd...).CombinedOutput()
	if err != nil {
		log.Printf("Error occured whilst executing az cli %s", string(out))
	}

	output := string(out)
	return output

}
