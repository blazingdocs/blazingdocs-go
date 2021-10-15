package parameters

type MergeParameters struct {
	DataSourceName string `json:"dataSourceName"`
	DataSourceType string `json:"dataSourceType"`
	Sequence       bool   `json:"sequence"`
	ParseColumns   bool   `json:"parseColumns"`
	Strict         bool   `json:"strict"`
}
