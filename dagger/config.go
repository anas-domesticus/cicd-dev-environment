package main

import (
	helper "ci-containers/go-build/ci-helper/types"
	"context"
	"dagger/cicd-dev-environment/internal/dagger"
	"encoding/json"
)

// LoadConfig populates the Config struct from a file
func LoadConfigFromFile(f *dagger.File) (*helper.Config, error) {
	fileContent, err := f.Contents(context.Background())
	if err != nil {
		return nil, err
	}
	cfg := helper.Config{}
	err = json.Unmarshal([]byte(fileContent), &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
