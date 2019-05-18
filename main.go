package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/abelal83/terraform_provider_cosmosdb/azcli"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/tidwall/gjson"
)

func main() {

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return azcli.Provider()
		},
	})

	//abutest()

	//c := azcli.NewClient()

	// output := c.AZCommand(cmd)

	// log.Print(output)
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

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return azcli.Provider()
		},
	})
}
