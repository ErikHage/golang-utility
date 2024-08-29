package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateUUID(t *testing.T) {
	id := GenerateUUID()
	length := len(id)

	assert.Truef(
		t,
		length == 36,
		"Expected uuid to be 36 characters, but was %d",
		length,
	)
}

func Test_Truncate(t *testing.T) {
	expected := "truncatedString"
	actual := Truncate("truncatedStringCutoff", 15)

	assert.Equalf(
		t,
		expected,
		actual,
		"Expected %s, but was %s",
		expected,
		actual,
	)
}

func Test_RandomString(t *testing.T) {
	actual := RandomString(15)

	assert.Truef(
		t,
		len(actual) == 15,
		"Expected string of length 15, but was %d [%s]",
		len(actual),
		actual,
	)
}
