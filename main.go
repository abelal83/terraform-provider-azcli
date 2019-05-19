package main

import (
	"fmt"
	"log"

	"github.com/abelal83/terraform_provider_cosmosdb/azcli"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/tidwall/gjson"
)

func main() {

	//cosmosDatabaseDelete()
	//cosmosDatabaseRead()
	//cosmosDatabaseCreate()

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return azcli.Provider()
		},
	})
}

func cosmosDatabaseRead() error {

	c := azcli.NewClient()
	name := "testdatabase"
	resourceGroupName := "terraform-provider"
	cosmosAccountName := "abx"
	cmd := []string{"cosmosdb", "database", "show", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	r, err := azcli.ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if !r.Found {
		// database doesn't exist
		log.Print("[INFO] database not found")
	}

	return nil
}

func cosmosDatabaseCreate() error {

	c := azcli.NewClient()
	name := "testdatabase"
	resourceGroupName := "terraform-provider"
	cosmosAccountName := "abx"

	cmd := []string{"cosmosdb", "database", "create", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	r, err := azcli.ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if r.AlreadyExists {
		// database already exists, lets just start to manage it.
		cmd = []string{"cosmosdb", "database", "show", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
		output = c.AZCommand(cmd)
		id := gjson.Get(output, "id")
		fmt.Print(id.Str)
		return nil
	}

	// new resource created
	id := gjson.Get(output, "id")
	fmt.Print(id.Str)
	return nil

}

func cosmosDatabaseDelete() error {
	c := azcli.NewClient()
	name := "testdatabase"
	resourceGroupName := "terraform-provider"
	cosmosAccountName := "abx"

	cmd := []string{"cosmosdb", "database", "delete", "--db-name", name, "-g", resourceGroupName, "-n", cosmosAccountName, "-o", "json"}
	output := c.AZCommand(cmd)

	r, err := azcli.ParseAzCliOutput(output)
	if err != nil {
		return err
	}

	if !r.Found || r.CliResponse == "" {
		// database doesn't exist
		log.Print("[INFO] database not found")
	}

	return nil
}
