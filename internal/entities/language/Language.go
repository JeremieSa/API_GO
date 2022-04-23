package entities

import "strings"

type Language struct {
	Code string
	Name string
}

type languageList []Language

func NewLanguage(code string, name string) Language {
	return Language{
		Code: code,
		Name: name,
	}
}

func (l languageList) Len() int {
	return len(l)
}

func (l languageList) Less(i, j int) bool {
	return strings.Compare(l[i].Code, l[j].Code) > 0
}

func (l languageList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func LanguageString(language Language) string {
	return "{" +
		"	Code: " + language.Code +
		"	Name: " + language.Name +
		"}"
}
