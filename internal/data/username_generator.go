package data

import (
	_ "embed"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
	"strings"
)

//go:embed embed/adjective_list.txt
var adjectiveList string

//go:embed embed/noun_list.txt
var nounList string

func GenerateUsername() string {
	adjectives := strings.Split(adjectiveList, "\n")
	nouns := strings.Split(nounList, "\n")

	randomAdjective := cases.Title(language.English).String(adjectives[rand.Intn(len(adjectives))])
	randomNoun := cases.Title(language.English).String(nouns[rand.Intn(len(nouns))])

	return randomAdjective + randomNoun
}
