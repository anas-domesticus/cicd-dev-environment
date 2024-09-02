package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type ImageComponents struct {
	Host      string
	ImageName string
	Tag       string
}

func parseImageReference(imageRef string) ImageComponents {
	var components ImageComponents

	parts := strings.Split(imageRef, "/")
	if len(parts) > 1 {
		components.Host = parts[0]
		imageRef = parts[1]
	}

	imageAndTag := strings.Split(imageRef, ":")
	components.ImageName = imageAndTag[0]
	if len(imageAndTag) > 1 {
		components.Tag = imageAndTag[1]
	} else {
		components.Tag = "latest" // Default tag
	}

	return components
}

func main() {
	// Define a flag for the image reference
	imageRefFlag := flag.String("image", "", "Image reference in the form [host]/[imageName]:[tag]")

	// Parse the input flags
	flag.Parse()

	if *imageRefFlag == "" {
		fmt.Println("Please provide an image reference using the -image flag")
		return
	}

	// Parse the image reference
	components := parseImageReference(*imageRefFlag)

	// Load the existing JSON file or create a new one if it doesn't exist.
	imgFilename := ".tilt/.images.json"
	data, err := os.ReadFile(imgFilename)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "error: could not read file %s: %v\n", imgFilename, err)
		os.Exit(1)
	}

	var jsonMap map[string]string
	if err = json.Unmarshal(data, &jsonMap); err != nil {
		jsonMap = make(map[string]string)
	}

	jsonMap[fmt.Sprintf("%s/%s", components.Host, components.ImageName)] = components.Tag

	data, err = json.MarshalIndent(jsonMap, "", "  ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(imgFilename, data, 0644)
	if err != nil {
		panic(err)
	}

	templateData, err := os.ReadFile(".tilt/templates/cd-overlay.tpl.yaml")
	if err != nil {
		panic(err)
	}
	// Create a new template and parse the YAML into it
	tmpl, err := template.New("test").Parse(string(templateData))
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, components)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fmt.Sprintf("./k8s/overlays/dev/cd/%s.yaml", components.ImageName), buffer.Bytes(), 0664)
	if err != nil {
		panic(err)
	}

}
