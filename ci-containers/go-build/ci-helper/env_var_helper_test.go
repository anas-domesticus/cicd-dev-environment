package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToEnvVars(t *testing.T) {
	jsonData := `{"ref":"refs/heads/main","before":"a1550e049d0c6e76773cbe396001b7f3d4f8ed83","after":"04f6f29ad4677c053ceca2b7e6f0287e726abdd5","compare_url":"http://localhost:3000/test-user/dev/compare/a1550e049d0c6e76773cbe396001b7f3d4f8ed83...04f6f29ad4677c053ceca2b7e6f0287e726abdd5","commits":[{"id":"04f6f29ad4677c053ceca2b7e6f0287e726abdd5","message":"update commit\n","url":"http://localhost:3000/test-user/dev/commit/04f6f29ad4677c053ceca2b7e6f0287e726abdd5","author":{"name":"Lewis Edginton","email":"lewis@example.com","username":""},"committer":{"name":"Lewis Edginton","email":"lewis@example.com","username":""},"verification":null,"timestamp":"2024-08-22T18:09:58+01:00","added":[],"removed":[],"modified":["ci-containers/go-build/ci-helper/main.go"]}],"head_commit":{"id":"04f6f29ad4677c053ceca2b7e6f0287e726abdd5","message":"update commit\n","url":"http://localhost:3000/test-user/dev/commit/04f6f29ad4677c053ceca2b7e6f0287e726abdd5","author":{"name":"Lewis Edginton","email":"lewis@example.com","username":""},"committer":{"name":"Lewis Edginton","email":"lewis@example.com","username":""},"verification":null,"timestamp":"2024-08-22T18:09:58+01:00","added":[],"removed":[],"modified":["ci-containers/go-build/ci-helper/main.go"]},"repository":{"id":1,"owner":{"id":1,"login":"test-user","full_name":"","email":"test@example.com","avatar_url":"http://localhost:3000/avatar/55502f40dc8b7c769880b10874abc9d0","language":"","is_admin":false,"last_login":"0001-01-01T00:00:00Z","created":"2024-08-22T16:52:09Z","restricted":false,"active":false,"prohibit_login":false,"location":"","website":"","description":"","visibility":"public","followers_count":0,"following_count":0,"starred_repos_count":0,"username":"test-user"},"name":"dev","full_name":"test-user/dev","description":"Dev Repo","empty":false,"private":false,"fork":false,"template":false,"parent":null,"mirror":false,"size":425,"html_url":"http://localhost:3000/test-user/dev","ssh_url":"git@localhost:test-user/dev.git","clone_url":"http://localhost:3000/test-user/dev.git","original_url":"","website":"","stars_count":0,"forks_count":0,"watchers_count":1,"open_issues_count":0,"open_pr_counter":0,"release_counter":0,"default_branch":"main","archived":false,"created_at":"2024-08-22T16:52:09Z","updated_at":"2024-08-22T17:09:10Z","permissions":{"admin":true,"push":true,"pull":true},"has_issues":true,"internal_tracker":{"enable_time_tracker":true,"allow_only_contributors_to_track_time":true,"enable_issue_dependencies":true},"has_wiki":true,"has_pull_requests":true,"has_projects":true,"ignore_whitespace_conflicts":false,"allow_merge_commits":true,"allow_rebase":true,"allow_rebase_explicit":true,"allow_squash_merge":true,"default_merge_style":"merge","avatar_url":"","internal":false,"mirror_interval":"","mirror_updated":"0001-01-01T00:00:00Z","repo_transfer":null},"pusher":{"id":1,"login":"test-user","full_name":"","email":"test@example.com","avatar_url":"http://localhost:3000/avatar/55502f40dc8b7c769880b10874abc9d0","language":"","is_admin":false,"last_login":"0001-01-01T00:00:00Z","created":"2024-08-22T16:52:09Z","restricted":false,"active":false,"prohibit_login":false,"location":"","website":"","description":"","visibility":"public","followers_count":0,"following_count":0,"starred_repos_count":0,"username":"test-user"},"sender":{"id":1,"login":"test-user","full_name":"","email":"test@example.com","avatar_url":"http://localhost:3000/avatar/55502f40dc8b7c769880b10874abc9d0","language":"","is_admin":false,"last_login":"0001-01-01T00:00:00Z","created":"2024-08-22T16:52:09Z","restricted":false,"active":false,"prohibit_login":false,"location":"","website":"","description":"","visibility":"public","followers_count":0,"following_count":0,"starred_repos_count":0,"username":"test-user"}}`

	var payload GiteaPayload
	err := json.Unmarshal([]byte(jsonData), &payload)
	assert.NoError(t, err)

	envVars := ToEnvVars(payload)

	assert.Equal(t, "refs/heads/main", envVars["GIT_REF"])
	assert.Equal(t, "a1550e049d0c6e76773cbe396001b7f3d4f8ed83", envVars["GIT_BEFORE"])
	assert.Equal(t, "04f6f29ad4677c053ceca2b7e6f0287e726abdd5", envVars["GIT_AFTER"])
	assert.Equal(t, "http://localhost:3000/test-user/dev/compare/a1550e049d0c6e76773cbe396001b7f3d4f8ed83...04f6f29ad4677c053ceca2b7e6f0287e726abdd5", envVars["GIT_COMPARE_URL"])
	assert.Equal(t, "1", envVars["GIT_REPOSITORY_ID"])
	assert.Equal(t, "test-user", envVars["GIT_REPOSITORY_OWNER_LOGIN"])
	assert.Equal(t, "dev", envVars["GIT_REPOSITORY_NAME"])
	assert.Equal(t, "true", envVars["GIT_REPOSITORY_PERMISSIONS_ADMIN"])
	assert.Equal(t, "2024-08-22T16:52:09Z", envVars["GIT_REPOSITORY_CREATED_AT"])
	assert.Equal(t, "2024-08-22T18:09:58+01:00", envVars["GIT_HEAD_COMMIT_TIMESTAMP"])
}
