package faker

import (
	"sync"
	"log"
	"fmt"
)

var bundle = NewBundle()

// NewBundle init fake bundle
func NewBundle() *Bundle {
	bundle := &Bundle{
		dataPath: dataPath(),
		backends: make(map[string]*Backend),
		lock:     &sync.RWMutex{},
	}

	bootBackends(bundle)
	bundle.SetLocale("en")

	return bundle
}

// async load backends
func bootBackends(bundle *Bundle) {
	wg := sync.WaitGroup{}

	files := dataFiles()
	for name, path := range files {
		wg.Add(1)

		go func(name, path string, bundle *Bundle) {
			defer wg.Done()
			backend, err := NewBackend(path)
			if err == nil {
				bundle.setBackend(name, backend)
			} else {
				log.Println(fmt.Sprintf("load backend[%s] err: %s", name, err))
			}
		}(name, path, bundle)
	}

	wg.Wait()

	for locale := range files {
		bundle.locales = append(bundle.locales, locale)
	}
}

func currentBackend() *Backend {
	return bundle.currentBackend
}

// Bundle support data
type Bundle struct {
	dataPath       string
	locales        []string
	currentLocale  string
	currentBackend *Backend
	backends       map[string]*Backend
	lock           *sync.RWMutex
}

func (b Bundle) setBackend(name string, backend *Backend) {
	b.lock.Lock()
	b.backends[name] = backend
	b.lock.Unlock()
}

// SetLocale current locale and backend
func (b *Bundle) SetLocale(locale string) *Bundle {
	if b.currentLocale == locale {
		return b
	}

	for _, l := range b.locales {
		if l == locale {
			b.currentLocale = l
		}
	}

	if backend, ok := b.backends[locale]; ok {
		b.currentBackend = backend
	}

	return b
}
