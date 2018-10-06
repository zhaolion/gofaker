package name

import (
	"github.com/hoisie/mustache"
	"github.com/zhaolion/faker/random"
)

// FakeName fake random name
type FakeName struct {
	MaleFirstNames    []string `yaml:"male_first_name,flow"`
	MaleMiddleNames   []string `yaml:"male_middle_name,flow"`
	MaleLastNames     []string `yaml:"male_last_name,flow"`
	FemaleFirstNames  []string `yaml:"female_first_name,flow"`
	FemaleMiddleNames []string `yaml:"female_middle_name,flow"`
	FemaleLastNames   []string `yaml:"female_last_name,flow"`
	FirstNames        []string `yaml:"first_name,flow"`
	LastNames         []string `yaml:"last_name,flow"`
	Prefixs           []string `yaml:"prefix,flow"`
	MalePrefixs       []string `yaml:"male_prefix,flow"`
	FemalePrefixs     []string `yaml:"female_prefix,flow"`
	Suffixs           []string `yaml:"suffix,flow"`
	Names             []string `yaml:"name,flow"`
	NameWithMiddles   []string `yaml:"name_with_middle,flow"`
}

// MaleFirstName male first name
func (fn *FakeName) MaleFirstName() string {
	return random.PickString(fn.MaleFirstNames)
}

// MaleMiddleName male middle name
func (fn *FakeName) MaleMiddleName() string {
	return random.PickString(fn.MaleMiddleNames)
}

// MaleLastName male last name
func (fn *FakeName) MaleLastName() string {
	return random.PickString(fn.MaleLastNames)
}

// FemaleFirstName female first name
func (fn *FakeName) FemaleFirstName() string {
	return random.PickString(fn.FemaleFirstNames)
}

// FemaleMiddleName female middle name
func (fn *FakeName) FemaleMiddleName() string {
	return random.PickString(fn.FemaleMiddleNames)
}

// FemaleFirstName female last name
func (fn *FakeName) FemaleLastName() string {
	return random.PickString(fn.FemaleLastNames)
}

// FirstName first name
func (fn *FakeName) FirstName() string {
	name := random.PickString(fn.FirstNames)
	return random.Numerify(mustache.Render(name, fn))
}

// MiddleName middle name
func (fn *FakeName) MiddleName() string {
	return fn.LastName()
}

// LastName last name
func (fn *FakeName) LastName() string {
	name := random.PickString(fn.LastNames)
	return random.Numerify(mustache.Render(name, fn))
}

// Prefix prefix of name
func (fn *FakeName) Prefix() string {
	return random.PickString(fn.Prefixs)
}

// MalePrefix prefix of male name
func (fn *FakeName) MalePrefix() string {
	return random.PickString(fn.MalePrefixs)
}

// FemalePrefix prefix of female name
func (fn *FakeName) FemalePrefix() string {
	return random.PickString(fn.FemalePrefixs)
}

// Suffix suffix of name
func (fn *FakeName) Suffix() string {
	return random.PickString(fn.Suffixs)
}

// Name random name
func (fn *FakeName) Name() string {
	name := random.PickString(fn.Names)
	return random.Numerify(mustache.Render(name, fn))
}

// NameWithMiddle middle format name
func (fn *FakeName) NameWithMiddle() string {
	name := random.PickString(fn.NameWithMiddles)
	return random.Numerify(mustache.Render(name, fn))
}
