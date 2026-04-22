package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeys(t *testing.T) {
	apiKey := "GR^$YDFY$EWYG#$&%$43ygdf%^gvdy$EWrdf"
	header := http.Header{
		"Authorization": []string{"ApiKey " + apiKey},
	}
	if key, err := GetAPIKey(header); err == nil {
		if key != apiKey {
			t.Fatalf("GetAPIKey returned a wrong API key: %s != %s", key, apiKey)
		}
	}
	if key, err := GetAPIKey(http.Header{}); err != nil {
		if key != "" || err != ErrNoAuthHeaderIncluded {
			t.Fatalf("GetAPIKey didn't handle no auth headers correctly")
		}
	}

}
