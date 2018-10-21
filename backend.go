package gofaker

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jinzhu/copier"
	"github.com/zhaolion/gofaker/address"
	"github.com/zhaolion/gofaker/name"
	"github.com/zhaolion/gofaker/phonenumber"
	"gopkg.in/yaml.v2"
)

// Backend of random base data
type Backend struct {
	Phone   *phonenumber.FakePhone `yaml:"fake_phone"`
	Name    *name.FakeName         `yaml:"fake_name"`
	Address *address.FakeAddress   `yaml:"fake_address"`
}

// NewLocaleBackend init backend from locale file
// en locale file contains default value
func NewLocaleBackend(locale string) (*Backend, error) {
	enBackend, err := initDefaultBackend()
	if err != nil {
		return nil, fmt.Errorf("can't load default backend err: %s", err)
	}

	return initBackend(localeFilePath(locale), *enBackend)
}

func initBackend(filename string, en Backend) (*Backend, error) {
	var target Backend
	if err := copier.Copy(&target, &en); err != nil {
		return nil, err
	}

	if err := unmarshalBackend(filename, &target); err != nil {
		return nil, err
	}

	return &target, nil
}

func initDefaultBackend() (*Backend, error) {
	backend := &Backend{}
	if err := unmarshalBackend(localeFilePath("en"), backend); err != nil {
		return nil, err
	}

	if backend.Name == nil {
		log.Fatal("default backend name should not be nil")
	}

	backend.Address.SetFakeName(backend.Name)

	return backend, nil
}

func unmarshalBackend(filename string, bak *Backend) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bs, bak)
	if err != nil {
		return err
	}

	return nil
}
