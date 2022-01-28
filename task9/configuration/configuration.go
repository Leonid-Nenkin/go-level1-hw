package configuration

import (
	"json"
	"os"
)

type Configuration struct {
	Port       string `json:"port"`
	DbUrl      string `json:"db_url"`
	JaegerUrl  string `json:"jaeger_url"`
	SomeAppId  string `json:"some_app_id"`
	SomeAppKey string `json:"some_app_key"`
}

func Load() (cfg *Configuration, err error) {

	contents, err := os.ReadFile("configuration.json")

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(contents, cfg); err != nil {
		panic(err)
	}

	return cfg, nil
}
