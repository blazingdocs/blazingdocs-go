package config

var Default = defaultInit()

type Config struct {
	ApiKey  string
	BaseURL string
}

func defaultInit() *Config {
	baseUrl := "https://api.blazingdocs.com"
	return &Config{"", baseUrl}
}

func Init(apiKey string) *Config {
	baseUrl := "https://api.blazingdocs.com"
	return &Config{apiKey, baseUrl}
}
