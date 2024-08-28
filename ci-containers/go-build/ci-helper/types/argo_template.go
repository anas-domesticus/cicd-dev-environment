package types

type ArgoTemplate struct {
	Name   string `json:"name"`
	Inputs struct {
		Parameters []ArgoParameter `json:"parameters"`
	} `json:"inputs"`
	Outputs struct {
	} `json:"outputs"`
	Metadata struct {
	} `json:"metadata"`
	Container struct {
		Name      string   `json:"name"`
		Image     string   `json:"image"`
		Command   []string `json:"command"`
		Resources struct {
		} `json:"resources"`
	} `json:"container"`
}

type ArgoParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (at *ArgoTemplate) ParametersToEnvVars() map[string]string {
	envVars := make(map[string]string)
	for _, param := range at.Inputs.Parameters {
		envVars[param.Name] = param.Value
	}
	return envVars
}

func (ap *ArgoParameter) IsPayload() bool {
	return ap.Name == "message"
}
