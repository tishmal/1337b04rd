package integration

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSessionCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-session", func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:  "session_id",
			Value: "some-session-id",
		})
		w.WriteHeader(http.StatusOK)
	})

	mux.HandleFunc("/check-session", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Session ID: %s", cookie.Value)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/set-session")
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	cookies := resp.Cookies()
	require.Len(t, cookies, 1)
	sessionCookie := cookies[0]

	require.Equal(t, "session_id", sessionCookie.Name)
	require.NotEmpty(t, sessionCookie.Value)

	req, err := http.NewRequest("GET", ts.URL+"/check-session", nil)
	require.NoError(t, err)
	req.AddCookie(sessionCookie)

	resp, err = http.DefaultClient.Do(req)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)
	require.Contains(t, string(body), "Session ID:")
}
