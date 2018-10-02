package faker

import (
	"os"
	"testing"

	"github.com/ZhaoLion/faker/random"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	random.Seed(1000)
	code := m.Run()
	os.Exit(code)
}

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
