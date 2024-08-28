package types

import (
	"encoding/json"
	"errors"
	"os"
)

// Config holds the environment variables
type Config struct {
	ArgoPodName                   string       `env:"ARGO_POD_NAME"`
	ArgoPodUID                    string       `env:"ARGO_POD_UID"`
	GoDebug                       string       `env:"GODEBUG"`
	ArgoWorkflowName              string       `env:"ARGO_WORKFLOW_NAME"`
	ArgoWorkflowUID               string       `env:"ARGO_WORKFLOW_UID"`
	ArgoContainerName             string       `env:"ARGO_CONTAINER_NAME"`
	ArgoTemplate                  ArgoTemplate `env:"ARGO_TEMPLATE"`
	ArgoNodeID                    string       `env:"ARGO_NODE_ID"`
	ArgoIncludeScriptOutput       string       `env:"ARGO_INCLUDE_SCRIPT_OUTPUT"`
	ArgoDeadline                  string       `env:"ARGO_DEADLINE"`
	ArgoProgressFile              string       `env:"ARGO_PROGRESS_FILE"`
	ArgoProgressPatchTickDuration string       `env:"ARGO_PROGRESS_PATCH_TICK_DURATION"`
	ArgoProgressFileTickDuration  string       `env:"ARGO_PROGRESS_FILE_TICK_DURATION"`
}

// LoadConfig populates the Config struct with environment variables
func LoadConfigFromEnvVars() (*Config, error) {
	argoTemplate := ArgoTemplate{}
	err := json.Unmarshal([]byte(os.Getenv("ARGO_TEMPLATE")), &argoTemplate)
	if err != nil {
		return nil, err
	}
	return &Config{
		ArgoPodName:                   os.Getenv("ARGO_POD_NAME"),
		ArgoPodUID:                    os.Getenv("ARGO_POD_UID"),
		GoDebug:                       os.Getenv("GODEBUG"),
		ArgoWorkflowName:              os.Getenv("ARGO_WORKFLOW_NAME"),
		ArgoWorkflowUID:               os.Getenv("ARGO_WORKFLOW_UID"),
		ArgoContainerName:             os.Getenv("ARGO_CONTAINER_NAME"),
		ArgoTemplate:                  argoTemplate,
		ArgoNodeID:                    os.Getenv("ARGO_NODE_ID"),
		ArgoIncludeScriptOutput:       os.Getenv("ARGO_INCLUDE_SCRIPT_OUTPUT"),
		ArgoDeadline:                  os.Getenv("ARGO_DEADLINE"),
		ArgoProgressFile:              os.Getenv("ARGO_PROGRESS_FILE"),
		ArgoProgressPatchTickDuration: os.Getenv("ARGO_PROGRESS_PATCH_TICK_DURATION"),
		ArgoProgressFileTickDuration:  os.Getenv("ARGO_PROGRESS_FILE_TICK_DURATION"),
	}, nil
}

func (c *Config) ExtractPayload() (*GiteaPayload, error) {
	for _, v := range c.ArgoTemplate.Inputs.Parameters {
		if v.IsPayload() {
			payload := &GiteaPayload{}
			err := json.Unmarshal([]byte(v.Value), payload)
			if err != nil {
				return nil, err
			}
			return payload, nil
		}
	}
	return nil, errors.New("cannot find Argo Template Message")
}
