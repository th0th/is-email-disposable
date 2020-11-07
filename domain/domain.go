package domain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	byteValue, err := ioutil.ReadAll(domainsJsonFile)

	if err != nil {
		return nil, fmt.Errorf("an error occurred while reading the domains file: %v", err)
	}

	var domainsSlice []string

	err = json.Unmarshal(byteValue, &domainsSlice)

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
