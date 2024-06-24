package data

import (
	_ "embed"
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

	randomAdjective := adjectives[rand.Intn(len(adjectives))]
	randomNoun := nouns[rand.Intn(len(nouns))]

	return strings.Title(randomAdjective) + strings.Title(randomNoun)
}
