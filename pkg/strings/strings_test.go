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

func Test_NullSafeString_NotNil(t *testing.T) {
	expected := "some string"
	actual := NullSafeString(&expected)

	assert.NotNilf(
		t,
		actual,
		"Expected string to not be nil, but was nil",
	)
}

func Test_NullSafeString_Nil(t *testing.T) {
	expected := ""
	actual := NullSafeString(&expected)

	assert.Equalf(
		t,
		expected,
		actual,
		"Expected string to be empty string, but was %s",
		actual,
	)
}
