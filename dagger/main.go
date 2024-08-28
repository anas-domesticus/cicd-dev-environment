package main

import (
	"context"
	"dagger/cicd-dev-environment/internal/dagger"
	"fmt"
	"math"
	"math/rand"
)

type CicdDevEnvironment struct{}

// Returns a container that echoes whatever string argument is provided
func (m *CicdDevEnvironment) ContainerEcho(stringArg string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg})
}

// CI Entrypoint
func (m *CicdDevEnvironment) Ci(source *dagger.Directory) (string, error) {
	cfg, err := LoadConfigFromFile(source.File("argo.json"))
	if err != nil {
		return "Load config error", err
	}
	payload, err := cfg.ExtractPayload()
	if err != nil {
		return "Extract payload error", err
	}

	fmt.Printf("Got %d commits...\n", len(payload.Commits))
	for i := range payload.Commits {
		fmt.Printf("- %s\n", payload.Commits[i].Message)
	}

	return m.Publish(context.Background(), source, "example")
}

// Return the result of running unit tests
func (m *CicdDevEnvironment) Test(ctx context.Context, source *dagger.Directory, app string) (string, error) {
	return m.BuildEnv(source).
		WithExec([]string{"go", "test", appToModule(app)}).
		Stdout(ctx)
}

// Returns an application container
func (m *CicdDevEnvironment) Build(source *dagger.Directory, app string) *dagger.Container {
	binary := m.BuildEnv(source).
		WithExec([]string{"go", "build", "-o", fmt.Sprintf("/%s", app), appToModule(app)}).
		File(fmt.Sprintf("/%s", app))

	return dag.Container().From("gcr.io/distroless/static-debian12").
		WithFile("/entrypoint", binary, dagger.ContainerWithFileOpts{
			Permissions: 0755,
			Owner:       "root",
		}).WithEntrypoint([]string{"/entrypoint"})
}

// Publish the application container after building and testing it on-the-fly
func (m *CicdDevEnvironment) Publish(ctx context.Context, source *dagger.Directory, app string) (string, error) {
	// call Dagger Function to run unit tests
	_, err := m.Test(ctx, source, app)
	if err != nil {
		return "", err
	}
	// call Dagger Function to build the application image
	// publish the image to ttl.sh
	address, err := m.Build(source, app).
		Publish(ctx, fmt.Sprintf("ttl.sh/example-%.0f", math.Floor(rand.Float64()*10000000))) //#nosec
	if err != nil {
		return "", err
	}
	return address, nil
}

// Build a ready-to-use development environment
func (m *CicdDevEnvironment) BuildEnv(source *dagger.Directory) *dagger.Container {
	return dag.Container().
		From("golang:1.22.5").
		WithEnvVariable("CGO_ENABLED", "0").
		WithEnvVariable("GOOS", "linux").
		WithEnvVariable("GOARCH", "amd64").
		WithDirectory("/src", source).
		WithWorkdir("/src")
}

func appToModule(s string) string {
	return fmt.Sprintf("./cmd/%s", s)
}
