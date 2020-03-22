package lib

import (
	"fmt"
	"strings"
)

type LangSpec struct {
	Greeting string
	NoName   string
}

func supportedLanguages() map[string]LangSpec {
	return map[string]LangSpec{
		"en": {"Hello", "everyone"},
		"fr": {"Bonjour", "tout le monde"},
	}
}

// Greetings takes a name as single argument and return a greeting for that name
func Greetings(language, name string) (string, error) {
	lang, found := supportedLanguages()[language]
	if !found {
		return "", fmt.Errorf("unknown language: '%s'", lang)
	}

	if name == "" {
		name = lang.NoName
	}

	msg := fmt.Sprintf("%s %s!", lang.Greeting, strings.Title(name))
	return msg, nil
}
