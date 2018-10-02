package faker

import (
	"testing"
)

func TestNewBundle(t *testing.T) {
	bundle := NewBundle()
	if len(bundle.backends) != len(bundle.locales) {
		t.Errorf("load bundle backends should contain size of %d", len(bundle.locales))
	}

	for name, backend := range bundle.backends {
		if backend == nil {
			t.Errorf("load bundle backends[%s] error", name)
		}
	}
}
