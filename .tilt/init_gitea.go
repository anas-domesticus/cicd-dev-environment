package main

import (
	"code.gitea.io/sdk/gitea"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

func main() {
	initGitea()
}

func initGitea() {
	giteaUser := "test-user"
	giteaPass := "password"
	giteaUrl := "http://localhost:3000"
	formData := url.Values{
		"user_name": {giteaUser},
		"email":     {"test@example.com"},
		"password":  {giteaPass},
		"retype":    {giteaPass},
	}
	resp, err := http.PostForm(fmt.Sprintf("%s/user/sign_up", giteaUrl), formData)
	if err != nil {
		fmt.Printf("Failed to post to gitea server, %s\n", err.Error())
		fmt.Println(resp)
		os.Exit(1)
	}
	if resp.StatusCode != 200 {
		fmt.Println(resp)
		os.Exit(1)
	}
	giteaClient, err := gitea.NewClient(giteaUrl)
	if err != nil {
		fmt.Printf("Failed to get gitea client, %s\n", err.Error())
		os.Exit(1)
	}
	giteaClient.SetBasicAuth(giteaUser, giteaPass)

	repo, _, err := giteaClient.CreateRepo(gitea.CreateRepoOption{
		Name:          "dev",
		Description:   "Dev Repo",
		DefaultBranch: "main",
		Private:       false,
	})
	if err != nil {
		fmt.Printf("Failed to get create repo, %s\n", err.Error())
		os.Exit(1)
	}

	_, _, err = giteaClient.CreateRepoHook(giteaUser, repo.Name, gitea.CreateHookOption{
		Type: "gitea",
		Config: map[string]string{
			"url":          "http://webhook-eventsource-svc.workflows.svc.cluster.local:12000/",
			"content_type": "json",
		},
		Events: []string{"push", "pr"},
		Active: true,
	})
	if err != nil {
		fmt.Printf("Failed to create webhook, %s\n", err.Error())
		os.Exit(1)
	}
}
