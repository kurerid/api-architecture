package initialize

import (
	"api-architecture/pkg/repository"
	"encoding/json"
	"github.com/execaus/exloggo"
	"os"
)

func Config(path string) (*repository.Config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		exloggo.Fatal(err.Error())
		return nil, err
	}

	var config repository.Config
	err = json.Unmarshal(b, &config)
	if err != nil {
		exloggo.Fatal(err.Error())
		return nil, err

	}

	return &config, nil
}
