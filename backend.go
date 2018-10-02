package faker

import (
	"io/ioutil"
	"os"

	"github.com/ZhaoLion/faker/phonenumber"
	"gopkg.in/yaml.v2"
)

// Backend of random base data
type Backend struct {
	Phone *phonenumber.FakePhone `yaml:"fake_phone"`
}

// NewBackend unmarshal backend from file
func NewBackend(filename string) (*Backend, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var backend Backend
	err = yaml.Unmarshal(bs, &backend)
	if err != nil {
		return nil, err
	}

	return &backend, nil
}
