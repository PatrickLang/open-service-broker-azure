package postgresql

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateIdentifier(t *testing.T) {
	regex, err := regexp.Compile(`^[a-z][a-z\d]{9}$`)
	assert.Nil(t, err)
	for range [100]struct{}{} {
		identifier := generateIdentifier()
		assert.True(t, regex.MatchString(identifier))
	}
}
