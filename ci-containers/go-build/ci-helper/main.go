package main

import (
	"ci-containers/go-build/ci-helper/types"
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"os"
	"os/exec"
)

func main() {
	err := cloneRepo("http://gitea.gitea.svc.cluster.local:3000/test-user/dev.git", "/src", nil)
	if err != nil {
		fmt.Printf("Error cloning repo: %v\n", err)
		os.Exit(1)
	}

	cfg, err := types.LoadConfigFromEnvVars()
	if err != nil {
		fmt.Printf("Load config error: %v\n", err)
		os.Exit(1)
	}

	if err = writeConfigToFile(cfg, "/src/argo.json"); err != nil {
		fmt.Printf("Error writing config to file: %v\n", err)
		os.Exit(1)
	}

	cmdName := "dagger"
	cmdArgs := []string{
		"call",
		"ci",
		"--source=/src/.",
	}

	if err := runSubcommand(cmdName, cmdArgs); err != nil {
		fmt.Println("Error:", err)
	}

}

func runSubcommand(cmdName string, cmdArgs []string) error {
	cmd := exec.Command(cmdName, cmdArgs...)

	env := os.Environ()
	cmd.Env = env
	cmd.Dir = "/src"

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}
	return nil
}

func cloneRepo(url, directory string, auth *http.BasicAuth) error {
	fmt.Printf("Cloning repository from %s to %s...\n", url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:      url,
		Progress: nil,
		Auth:     auth,
	})

	if err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	fmt.Println("Repository cloned successfully!")
	return nil
}

func writeConfigToFile(cfg interface{}, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("unable to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cfg); err != nil {
		return fmt.Errorf("unable to encode config to JSON: %w", err)
	}
	return nil
}
