package strings

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

func GenerateUUID() string {
	id := uuid.New()
	return id.String()
}

func Truncate(str string, length int) string {
	if len(str) <= length {
		return str
	}

	return str[:length]
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NullSafeString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}

func NullSafeBool(b *bool) string {
	if b == nil {
		return ""
	}

	return strconv.FormatBool(*b)
}

func NullSafeTime(t *time.Time) string {
	if t == nil {
		return ""
	}

	return t.String()
}

func NullSafeInt(i *int) string {
	if i == nil {
		return ""
	}

	return strconv.Itoa(*i)
}

func NullSafeFloat32(f *float32) string {
	if f == nil {
		return ""
	}

	return fmt.Sprintf("%f", *f)
}
