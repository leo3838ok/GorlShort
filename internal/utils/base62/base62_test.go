package base62

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	want := "g7"
	got := Encode(999)

	assert.Equal(t, want, got)
}

func TestDecode(t *testing.T) {
	want := 999
	got := Decode("g7")

	assert.Equal(t, want, got)
}
