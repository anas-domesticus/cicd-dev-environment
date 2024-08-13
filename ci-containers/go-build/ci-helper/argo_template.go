package main

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

func (ap *ArgoParameter) IsPayload() bool {
	return ap.Name == "message"
}
