package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/tidwall/gjson"
)

func main() {

	// plugin.Serve(&plugin.ServeOpts{
	// 	ProviderFunc: func() terraform.ResourceProvider {
	// 		return Provider()
	// 	},
	// })

	abutest()

	c := NewClient("fff")
	cmd := []string{"account", "show"}
	output := c.AZCommand(cmd)
	log.Print(output)
}

func abutest() {

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
}
