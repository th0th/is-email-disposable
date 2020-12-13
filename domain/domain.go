package domain

import (
	"encoding/json"
	"fmt"
	"os"
)

type Domain struct {
	domains map[string]bool
}

func NewDomain() (*Domain, error) {
	var err error

	domainsJsonFile, err := os.Open("domains.json")

	if err != nil {
		return nil, fmt.Errorf("an error occurred while opening the domains file: %v", err)
	}

	var domainsSlice []string

	err = json.NewDecoder(domainsJsonFile).Decode(&domainsSlice)

	if err != nil {
		return nil, fmt.Errorf("an error occurred while parsing the domains file: %v", err)
	}

	domains := make(map[string]bool)

	for _, domain := range domainsSlice {
		domains[domain] = true
	}

	return &Domain{
		domains: domains,
	}, nil
}

func (d *Domain) GetIsDisposable(emailOrDomain string) bool {
	return d.domains[emailOrDomain]
}
