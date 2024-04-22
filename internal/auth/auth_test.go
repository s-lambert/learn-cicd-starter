package auth

import (
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := make(map[string][]string)
	h["Authorization"] = []string{"ApiKey myApiKey"}

	key, err := GetAPIKey(h)

	if err != nil {
		t.Fatal("Got an error, expected myApiKey")
	}
	if !reflect.DeepEqual(key, "myApiKey") {
		t.Fatalf("Got %s, expected myApiKey", key)
	}
}
