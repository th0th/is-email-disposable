package common

type DomainService interface {
	Check(emailOrDomain string) *DomainServiceCheckResult
}

type DomainServiceCheckResult struct {
	IsDisposable bool `json:"isDisposable"`
}
