package pull_request

import (
	"crypto/x509"
	"encoding/pem"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func writeMRListResponse(t *testing.T, w io.Writer) {
	t.Helper()
	f, err := os.Open("fixtures/gitlab_mr_list_response.json")
	require.NoErrorf(t, err, "error opening fixture file: %v", err)

	_, err = io.Copy(w, f)
	require.NoErrorf(t, err, "error writing response: %v", err)
}

func TestGitLabServiceCustomBaseURL(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, path+"?per_page=100", r.URL.RequestURI())
		writeMRListResponse(t, w)
	})

	svc, err := NewGitLabService("", server.URL, "278964", nil, "", "", false, nil)
	require.NoError(t, err)

	_, err = svc.List(t.Context())
	require.NoError(t, err)
}

func TestGitLabServiceToken(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "token-123", r.Header.Get("Private-Token"))
		writeMRListResponse(t, w)
	})

	svc, err := NewGitLabService("token-123", server.URL, "278964", nil, "", "", false, nil)
	require.NoError(t, err)

	_, err = svc.List(t.Context())
	require.NoError(t, err)
}

func TestList(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, path+"?per_page=100", r.URL.RequestURI())
		writeMRListResponse(t, w)
	})

	svc, err := NewGitLabService("", server.URL, "278964", []string{}, "", "", false, nil)
	require.NoError(t, err)

	prs, err := svc.List(t.Context())
	require.NoError(t, err)
	assert.Len(t, prs, 1)
	assert.Equal(t, 15442, prs[0].Number)
	assert.Equal(t, "Draft: Use structured logging for DB load balancer", prs[0].Title)
	assert.Equal(t, "use-structured-logging-for-db-load-balancer", prs[0].Branch)
	assert.Equal(t, "master", prs[0].TargetBranch)
	assert.Equal(t, "2fc4e8b972ff3208ec63b6143e34ad67ff343ad7", prs[0].HeadSHA)
	assert.Equal(t, "hfyngvason", prs[0].Author)
}

func TestListWithLabels(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, path+"?labels=feature%2Cready&per_page=100", r.URL.RequestURI())
		writeMRListResponse(t, w)
	})

	svc, err := NewGitLabService("", server.URL, "278964", []string{"feature", "ready"}, "", "", false, nil)
	require.NoError(t, err)

	_, err = svc.List(t.Context())
	require.NoError(t, err)
}

func TestListWithState(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	path := "/api/v4/projects/278964/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, path+"?per_page=100&state=opened", r.URL.RequestURI())
		writeMRListResponse(t, w)
	})

	svc, err := NewGitLabService("", server.URL, "278964", []string{}, "opened", "", false, nil)
	require.NoError(t, err)

	_, err = svc.List(t.Context())
	require.NoError(t, err)
}

func TestListWithStateTLS(t *testing.T) {
	tests := []struct {
		name        string
		tlsInsecure bool
		passCerts   bool
		requireErr  bool
	}{
		{
			name:        "TLS Insecure: true, No Certs",
			tlsInsecure: true,
			passCerts:   false,
			requireErr:  false,
		},
		{
			name:        "TLS Insecure: true, With Certs",
			tlsInsecure: true,
			passCerts:   true,
			requireErr:  false,
		},
		{
			name:        "TLS Insecure: false, With Certs",
			tlsInsecure: false,
			passCerts:   true,
			requireErr:  false,
		},
		{
			name:        "TLS Insecure: false, No Certs",
			tlsInsecure: false,
			passCerts:   false,
			requireErr:  true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				writeMRListResponse(t, w)
			}))
			defer ts.Close()

			var certs []byte
			if test.passCerts {
				for _, cert := range ts.TLS.Certificates {
					for _, c := range cert.Certificate {
						parsedCert, err := x509.ParseCertificate(c)
						require.NoError(t, err, "Failed to parse certificate")
						certs = append(certs, pem.EncodeToMemory(&pem.Block{
							Type:  "CERTIFICATE",
							Bytes: parsedCert.Raw,
						})...)
					}
				}
			}

			svc, err := NewGitLabService("", ts.URL, "278964", []string{}, "opened", "", test.tlsInsecure, certs)
			require.NoError(t, err)

			_, err = svc.List(t.Context())
			if test.requireErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGitLabListReturnsRepositoryNotFoundError(t *testing.T) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	path := "/api/v4/projects/nonexistent/merge_requests"

	mux.HandleFunc(path, func(w http.ResponseWriter, _ *http.Request) {
		// Return 404 status to simulate repository not found
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"message": "404 Project Not Found"}`))
	})

	svc, err := NewGitLabService("", server.URL, "nonexistent", []string{}, "", "", false, nil)
	require.NoError(t, err)

	prs, err := svc.List(t.Context())

	// Should return empty pull requests list
	assert.Empty(t, prs)

	// Should return RepositoryNotFoundError
	require.Error(t, err)
	assert.True(t, IsRepositoryNotFoundError(err), "Expected RepositoryNotFoundError but got: %v", err)
}
