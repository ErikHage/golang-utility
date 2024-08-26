package strings

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateUUID(t *testing.T) {
	id := uuid.New().String()
	length := len(id)

	assert.Truef(
		t,
		length == 36,
		"Expected uuid to be 36 characters, but was %d",
		length,
	)
}
