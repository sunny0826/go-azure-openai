package openai

import (
	"net/http"
)

const (
	apiVersion = "2022-12-01"
)

// ClientConfig is a configuration of a client.
type ClientConfig struct {
	apiKey       string
	HTTPClient   *http.Client
	BaseURL      string
	resourceName string
}

func DefaultConfig(apiKey, resourceName string) ClientConfig {
	return ClientConfig{
		HTTPClient:   &http.Client{},
		BaseURL:      "openai.azure.com/openai/deployments/",
		apiKey:       apiKey,
		resourceName: resourceName,
	}
}
