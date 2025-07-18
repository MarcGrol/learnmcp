package config

import (
	"flag"
)

type Config struct {
	UseSSE        bool
	UseStreamable bool
	Port          string
	BaseURL       string
}

func LoadConfig() Config {
	useSSE := flag.Bool("sse", false, "Use SSE transport instead of stdio")
	useStreamable := flag.Bool("http", false, "Use Streamable HTTP transport (easier for testing)")
	port := flag.String("port", "8080", "Port for SSE server")
	baseURL := flag.String("baseurl", "http://localhost", "Base URL for SSE server")

	flag.Parse()

	return Config{
		UseSSE:        *useSSE,
		UseStreamable: *useStreamable,
		Port:          *port,
		BaseURL:       *baseURL,
	}
}
