package main

import (
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"log"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Loading config...")
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	fmt.Println("Payload:")

	payload, err := cfg.ExtractPayload()
	if err != nil {
		fmt.Printf("Error extracting payload: %v", err)
		os.Exit(1)
	}

	envVars := ToEnvVars(*payload)

	cmdName := "ls"
	cmdArgs := []string{
		"-la",
		"/src",
	}

	fmt.Println("foo")

	err = cloneRepo("http://gitea.gitea.svc.cluster.local:3000/test-user/dev.git", "/src", nil)
	if err != nil {
		fmt.Printf("Error cloning repo: %v", err)
		os.Exit(1)
	}

	if err := runSubcommand(cmdName, cmdArgs, envVars); err != nil {
		fmt.Println("Error:", err)
	}

}

func runSubcommand(cmdName string, cmdArgs []string, envVars map[string]string) error {
	cmd := exec.Command(cmdName, cmdArgs...)

	env := os.Environ()
	for key, value := range envVars {
		env = append(env, fmt.Sprintf("%s=%s", key, value))
	}
	cmd.Env = env

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
