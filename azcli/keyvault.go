package azcli

import (
	"log"

	"github.com/tidwall/gjson"
)

func keyvaultCreate(resourceGroup, name string, client Client) (string, error) {

	cmd := []string{
		"keyvault", "create",
		"--resource-group", resourceGroup,
		"--name", name,
	}

	log.Printf("Formatted command %s", cmd)

	output := client.AZCommand(cmd)

	result, err := ParseAzCliOutput(output)
	if err != nil {
		return "", err
	}

	if result.AlreadyExists {

		output, err = keyvaultRead(resourceGroup, name, client)
		if err != nil {
			return "", err
		}
	}

	id := gjson.Get(output, "id")
	return id.Str, nil
}

func keyvaultRead(resourceGroup, name string, client Client) (string, error) {

	cmd := []string{
		"keyvault", "show",
		"--resource-group", resourceGroup,
		"--name", name,
	}

	log.Printf("Formatted command %s", cmd)

	id := gjson.Get(output, "id")
	return id.Str, nil
}

func keyvaultupdate(resourceGroup, name string) {

}

func keyvaultdelete(resourceGroup, name string) {

	cmd := []string{
		"keyvault", "delete",
		"--resource-group", resourceGroup,
		"--name", name,
	}

	log.Printf("Formatted command %s", cmd)
}
