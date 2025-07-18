package pull_request

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"code.gitea.io/sdk/gitea"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func giteaMockHandler(t *testing.T) func(http.ResponseWriter, *http.Request) {
	t.Helper()
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(r.RequestURI)
		switch r.RequestURI {
		case "/api/v1/version":
			_, err := io.WriteString(w, `{"version":"1.17.0+dev-452-g1f0541780"}`)
			if err != nil {
				t.Fail()
			}
		case "/api/v1/repos/test-argocd/pr-test/pulls?limit=0&page=1&state=open":
			_, err := io.WriteString(w, `[{
				"id": 50721,
				"url": "https://gitea.com/test-argocd/pr-test/pulls/1",
				"number": 1,
				"user": {
					"id": 4476,
					"login": "graytshirt",
					"full_name": "Dan",
					"email": "graytshirt@noreply.gitea.io",
					"avatar_url": "https://secure.gravatar.com/avatar/2446c67bcd59d71f6ae3cf376ec2ae37?d=identicon",
					"language": "",
					"is_admin": false,
					"last_login": "0001-01-01T00:00:00Z",
					"created": "2020-04-07T01:14:36+08:00",
					"restricted": false,
					"active": false,
					"prohibit_login": false,
					"location": "",
					"website": "",
					"description": "",
					"visibility": "public",
					"followers_count": 0,
					"following_count": 4,
					"starred_repos_count": 1,
					"username": "graytshirt"
				},
				"title": "add an empty file",
				"body": "",
				"labels": [{"id": 1, "name": "label1", "color": "00aabb", "description": "foo", "url": ""}],
				"milestone": null,
				"assignee": null,
				"assignees": null,
				"state": "open",
				"is_locked": false,
				"comments": 0,
				"html_url": "https://gitea.com/test-argocd/pr-test/pulls/1",
				"diff_url": "https://gitea.com/test-argocd/pr-test/pulls/1.diff",
				"patch_url": "https://gitea.com/test-argocd/pr-test/pulls/1.patch",
				"mergeable": true,
				"merged": false,
				"merged_at": null,
				"merge_commit_sha": null,
				"merged_by": null,
				"base": {
					"label": "main",
					"ref": "main",
					"sha": "72687815ccba81ef014a96201cc2e846a68789d8",
					"repo_id": 21618,
					"repo": {
						"id": 21618,
						"owner": {
							"id": 31480,
							"login": "test-argocd",
							"full_name": "",
							"email": "",
							"avatar_url": "https://gitea.com/avatars/22d1b1d3f61abf95951c4a958731d848",
							"language": "",
							"is_admin": false,
							"last_login": "0001-01-01T00:00:00Z",
							"created": "2022-04-06T02:28:06+08:00",
							"restricted": false,
							"active": false,
							"prohibit_login": false,
							"location": "",
							"website": "",
							"description": "",
							"visibility": "public",
							"followers_count": 0,
							"following_count": 0,
							"starred_repos_count": 0,
							"username": "test-argocd"
						},
						"name": "pr-test",
						"full_name": "test-argocd/pr-test",
						"description": "",
						"empty": false,
						"private": false,
						"fork": false,
						"template": false,
						"parent": null,
						"mirror": false,
						"size": 28,
						"language": "",
						"languages_url": "https://gitea.com/api/v1/repos/test-argocd/pr-test/languages",
						"html_url": "https://gitea.com/test-argocd/pr-test",
						"ssh_url": "git@gitea.com:test-argocd/pr-test.git",
						"clone_url": "https://gitea.com/test-argocd/pr-test.git",
						"original_url": "",
						"website": "",
						"stars_count": 0,
						"forks_count": 0,
						"watchers_count": 1,
						"open_issues_count": 0,
						"open_pr_counter": 1,
						"release_counter": 0,
						"default_branch": "main",
						"archived": false,
						"created_at": "2022-04-06T02:32:09+08:00",
						"updated_at": "2022-04-06T02:33:12+08:00",
						"permissions": {
							"admin": false,
							"push": false,
							"pull": true
						},
						"has_issues": true,
						"internal_tracker": {
							"enable_time_tracker": true,
							"allow_only_contributors_to_track_time": true,
							"enable_issue_dependencies": true
						},
						"has_wiki": true,
						"has_pull_requests": true,
						"has_projects": true,
						"ignore_whitespace_conflicts": false,
						"allow_merge_commits": true,
						"allow_rebase": true,
						"allow_rebase_explicit": true,
						"allow_squash_merge": true,
						"default_merge_style": "merge",
						"avatar_url": "",
						"internal": false,
						"mirror_interval": "",
						"mirror_updated": "0001-01-01T00:00:00Z",
						"repo_transfer": null
					}
				},
				"head": {
					"label": "test",
					"ref": "test",
					"sha": "7bbaf62d92ddfafd9cc8b340c619abaec32bc09f",
					"repo_id": 21618,
					"repo": {
						"id": 21618,
						"owner": {
							"id": 31480,
							"login": "test-argocd",
							"full_name": "",
							"email": "",
							"avatar_url": "https://gitea.com/avatars/22d1b1d3f61abf95951c4a958731d848",
							"language": "",
							"is_admin": false,
							"last_login": "0001-01-01T00:00:00Z",
							"created": "2022-04-06T02:28:06+08:00",
							"restricted": false,
							"active": false,
							"prohibit_login": false,
							"location": "",
							"website": "",
							"description": "",
							"visibility": "public",
							"followers_count": 0,
							"following_count": 0,
							"starred_repos_count": 0,
							"username": "test-argocd"
						},
						"name": "pr-test",
						"full_name": "test-argocd/pr-test",
						"description": "",
						"empty": false,
						"private": false,
						"fork": false,
						"template": false,
						"parent": null,
						"mirror": false,
						"size": 28,
						"language": "",
						"languages_url": "https://gitea.com/api/v1/repos/test-argocd/pr-test/languages",
						"html_url": "https://gitea.com/test-argocd/pr-test",
						"ssh_url": "git@gitea.com:test-argocd/pr-test.git",
						"clone_url": "https://gitea.com/test-argocd/pr-test.git",
						"original_url": "",
						"website": "",
						"stars_count": 0,
						"forks_count": 0,
						"watchers_count": 1,
						"open_issues_count": 0,
						"open_pr_counter": 1,
						"release_counter": 0,
						"default_branch": "main",
						"archived": false,
						"created_at": "2022-04-06T02:32:09+08:00",
						"updated_at": "2022-04-06T02:33:12+08:00",
						"permissions": {
							"admin": false,
							"push": false,
							"pull": true
						},
						"has_issues": true,
						"internal_tracker": {
							"enable_time_tracker": true,
							"allow_only_contributors_to_track_time": true,
							"enable_issue_dependencies": true
						},
						"has_wiki": true,
						"has_pull_requests": true,
						"has_projects": true,
						"ignore_whitespace_conflicts": false,
						"allow_merge_commits": true,
						"allow_rebase": true,
						"allow_rebase_explicit": true,
						"allow_squash_merge": true,
						"default_merge_style": "merge",
						"avatar_url": "",
						"internal": false,
						"mirror_interval": "",
						"mirror_updated": "0001-01-01T00:00:00Z",
						"repo_transfer": null
					}
				},
				"merge_base": "72687815ccba81ef014a96201cc2e846a68789d8",
				"due_date": null,
				"created_at": "2022-04-06T02:34:24+08:00",
				"updated_at": "2022-04-06T02:34:24+08:00",
				"closed_at": null
			}]`)
			if err != nil {
				t.Fail()
			}
		}
	}
}

func TestGiteaContainLabels(t *testing.T) {
	cases := []struct {
		Name       string
		Labels     []string
		PullLabels []*gitea.Label
		Expect     bool
	}{
		{
			Name:   "Match labels",
			Labels: []string{"label1", "label2"},
			PullLabels: []*gitea.Label{
				{Name: "label1"},
				{Name: "label2"},
				{Name: "label3"},
			},
			Expect: true,
		},
		{
			Name:   "Not match labels",
			Labels: []string{"label1", "label4"},
			PullLabels: []*gitea.Label{
				{Name: "label1"},
				{Name: "label2"},
				{Name: "label3"},
			},
			Expect: false,
		},
		{
			Name:   "No specify",
			Labels: []string{},
			PullLabels: []*gitea.Label{
				{Name: "label1"},
				{Name: "label2"},
				{Name: "label3"},
			},
			Expect: true,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if got := giteaContainLabels(c.Labels, c.PullLabels); got != c.Expect {
				t.Errorf("expect: %v, got: %v", c.Expect, got)
			}
		})
	}
}

func TestGiteaList(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		giteaMockHandler(t)(w, r)
	}))
	host, err := NewGiteaService("", ts.URL, "test-argocd", "pr-test", []string{"label1"}, false)
	require.NoError(t, err)
	prs, err := host.List(t.Context())
	require.NoError(t, err)
	assert.Len(t, prs, 1)
	assert.Equal(t, 1, prs[0].Number)
	assert.Equal(t, "add an empty file", prs[0].Title)
	assert.Equal(t, "test", prs[0].Branch)
	assert.Equal(t, "main", prs[0].TargetBranch)
	assert.Equal(t, "7bbaf62d92ddfafd9cc8b340c619abaec32bc09f", prs[0].HeadSHA)
	assert.Equal(t, "graytshirt", prs[0].Author)
}

func TestGetGiteaPRLabelNames(t *testing.T) {
	Tests := []struct {
		Name           string
		PullLabels     []*gitea.Label
		ExpectedResult []string
	}{
		{
			Name: "PR has labels",
			PullLabels: []*gitea.Label{
				{Name: "label1"},
				{Name: "label2"},
				{Name: "label3"},
			},
			ExpectedResult: []string{"label1", "label2", "label3"},
		},
		{
			Name:           "PR does not have labels",
			PullLabels:     []*gitea.Label{},
			ExpectedResult: nil,
		},
	}
	for _, test := range Tests {
		t.Run(test.Name, func(t *testing.T) {
			labels := getGiteaPRLabelNames(test.PullLabels)
			assert.Equal(t, test.ExpectedResult, labels)
		})
	}
}

func TestGiteaListReturnsRepositoryNotFoundError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	// Handle version endpoint that Gitea client calls first
	mux.HandleFunc("/api/v1/version", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"version":"1.17.0+dev-452-g1f0541780"}`))
	})

	path := "/api/v1/repos/nonexistent/nonexistent/pulls?limit=0&page=1&state=open"

	mux.HandleFunc(path, func(w http.ResponseWriter, _ *http.Request) {
		// Return 404 status to simulate repository not found
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "404 Project Not Found"}`))
	})

	svc, err := NewGiteaService("", server.URL, "nonexistent", "nonexistent", []string{}, false)
	require.NoError(t, err)

	prs, err := svc.List(t.Context())

	// Should return empty pull requests list
	assert.Empty(t, prs)

	// Should return RepositoryNotFoundError
	require.Error(t, err)
	assert.True(t, IsRepositoryNotFoundError(err), "Expected RepositoryNotFoundError but got: %v", err)
}
