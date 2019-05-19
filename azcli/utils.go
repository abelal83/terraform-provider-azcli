package azcli

import (
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/gjson"
)

// ResourceState checks for state
type ResourceState struct {
	Found         bool
	AlreadyExists bool
	CliResponse   string
}

// ParseAzCliOutput parses az cli output
func ParseAzCliOutput(o string) (*ResourceState, error) {

	s := ResourceState{}

	switch gjson.Valid(o) {
	case true:
		log.Print("az cli output is valid json")
		s.Found = true
		return &s, nil
	case false:
		if strings.Contains(o, "Operation Failed: Resource Not Found") {
			s.Found = false
			return &s, nil
		}
		if strings.Contains(o, "Operation Failed: Resource Already Exists") {
			s.AlreadyExists = true
			return &s, nil
		}
		if strings.Compare(o, "") == 0 {
			// delete operations from az cli respond back with no output
			s.CliResponse = ""
			return &s, nil
		}
	}

	return &s, fmt.Errorf("Unhandled error meesage returned by AZ cli: %s", o)
}
