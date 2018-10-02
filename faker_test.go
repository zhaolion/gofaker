package faker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLocale(t *testing.T) {
	SetLocale("zh-CN")

	if bundle.currentLocale != "zh-CN" {
		t.Error("func SetLocale error")
	}

	backend, ok := bundle.backends["zh-CN"]
	if !ok {
		t.Fatal("func SetLocale backend error")
	}

	assert.Equal(t, backend, bundle.currentBackend)
}
