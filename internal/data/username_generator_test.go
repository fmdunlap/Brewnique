package data

import (
	"brewnique.fdunlap.com/internal/validator"
	"regexp"
	"testing"
)

func Test_GenerateUserName(t *testing.T) {
	NameRegex := regexp.MustCompile("^[a-zA-Z0-9]+$")

	t.Run("should generate a random name", func(t *testing.T) {
		name := GenerateUsername()
		if len(name) == 0 {
			t.Errorf("name should not be empty")
		}
		if !validator.Matches(name, NameRegex) {
			t.Errorf("name should match regex %s", NameRegex)
		}
	})
}
