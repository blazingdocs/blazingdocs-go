package models

type AccountModel struct {
	Id             int
	Name           string
	ApiKey         string
	ObsoleteApiKey string
	CreateadAt     string
	LastSyncedAt   string
	UpdatedAt      string
	IsDisabled     string
}
