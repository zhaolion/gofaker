package name

import (
	"github.com/hoisie/mustache"
	"github.com/zhaolion/faker/random"
)

// FakeName fake random name
type FakeName struct {
	MaleFirstNames          []string `yaml:"male_first_name,flow"`
	MaleMiddleNames         []string `yaml:"male_middle_name,flow"`
	MaleLastNames           []string `yaml:"male_last_name,flow"`
	FemaleFirstNames        []string `yaml:"female_first_name,flow"`
	FemaleMiddleNames       []string `yaml:"female_middle_name,flow"`
	FemaleLastNames         []string `yaml:"female_last_name,flow"`
	FirstNames              []string `yaml:"first_name,flow"`
	LastNames               []string `yaml:"last_name,flow"`
	Prefixs                 []string `yaml:"prefix,flow"`
	MalePrefixs             []string `yaml:"male_prefix,flow"`
	FemalePrefixs           []string `yaml:"female_prefix,flow"`
	Suffixs                 []string `yaml:"suffix,flow"`
	Names                   []string `yaml:"name,flow"`
	NameWithMiddles         []string `yaml:"name_with_middle,flow"`
	NobilityTitlePrefixs    []string `yaml:"nobility_title_prefix,flow"`
	NobilityTitles          []string `yaml:"nobility_title,flow"`
	OckerFirstNames         []string `yaml:"ocker_first_name,flow"`
	MalayMaleFirstNames     []string `yaml:"malay_male_first_name,flow"`
	MalayFemaleFirstNames   []string `yaml:"malay_female_first_name,flow"`
	ChineseMaleFirstNames   []string `yaml:"chinese_male_first_name,flow"`
	ChineseMaleLastNames    []string `yaml:"chinese_male_last_name,flow"`
	ChineseFemaleFirstNames []string `yaml:"chinese_female_first_name,flow"`
	MaleEnglishNames        []string `yaml:"male_english_name,flow"`
	FemaleEnglishNames      []string `yaml:"female_english_name,flow"`
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

// FemaleLastName female last name
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

// NobilityTitlePrefix  nobility_title_prefix
func (fn *FakeName) NobilityTitlePrefix() string {
	return random.PickString(fn.NobilityTitlePrefixs)
}

// NobilityTitle  nobility_title
func (fn *FakeName) NobilityTitle() string {
	return random.PickString(fn.NobilityTitles)
}

// OckerFirstName  ocker_first_name
func (fn *FakeName) OckerFirstName() string {
	return random.PickString(fn.OckerFirstNames)
}

// MalayMaleFirstName  malay_male_first_name
func (fn *FakeName) MalayMaleFirstName() string {
	return random.PickString(fn.MalayMaleFirstNames)
}

// MalayFemaleFirstName  malay_female_first_name
func (fn *FakeName) MalayFemaleFirstName() string {
	return random.PickString(fn.MalayFemaleFirstNames)
}

// ChineseMaleFirstName  chinese_male_first_name
func (fn *FakeName) ChineseMaleFirstName() string {
	return random.PickString(fn.ChineseMaleFirstNames)
}

// ChineseMaleLastName  chinese_male_last_name
func (fn *FakeName) ChineseMaleLastName() string {
	return random.PickString(fn.ChineseMaleLastNames)
}

// ChineseFemaleFirstName  chinese_female_first_name
func (fn *FakeName) ChineseFemaleFirstName() string {
	return random.PickString(fn.ChineseFemaleFirstNames)
}

// MaleEnglishName  male_english_name
func (fn *FakeName) MaleEnglishName() string {
	return random.PickString(fn.MaleEnglishNames)
}

// FemaleEnglishName  female_english_name
func (fn *FakeName) FemaleEnglishName() string {
	return random.PickString(fn.FemaleEnglishNames)
}
