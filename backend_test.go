package faker

import (
	"testing"
)

func TestNewBackend(t *testing.T) {
	backend, err := NewBackend(dataFilePath("en.yml"))
	if err != nil {
		t.Fatalf("NewBackend with err: %s", err)
	}
	if backend == nil {
		t.Fatalf("NewBackend return nil")
	}
}
