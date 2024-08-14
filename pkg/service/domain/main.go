package domain

import (
	_ "embed"
	"encoding/csv"
	"strings"

	"github.com/go-errors/errors"
	"github.com/th0th/is-email-disposable/pkg/common"
)

type service struct {
	domainsMap map[string]bool
}

func New() (common.DomainService, error) {
	csvReader := csv.NewReader(strings.NewReader(domainsCsvFile))
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, 0)
	}

	domainsMap := map[string]bool{}
	for i := range records {
		domainsMap[records[i][0]] = true
	}

	return &service{
		domainsMap: domainsMap,
	}, nil
}

func (s *service) Check(emailOrDomain string) *common.DomainServiceCheckResult {
	domain := emailOrDomain

	atIndex := strings.LastIndex(emailOrDomain, "@")
	if atIndex > -1 {
		domain = emailOrDomain[atIndex+1:]
	}

	return &common.DomainServiceCheckResult{
		IsDisposable: s.domainsMap[domain],
	}
}

//go:embed domains.csv
var domainsCsvFile string
