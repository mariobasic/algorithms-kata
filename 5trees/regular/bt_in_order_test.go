package regular

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_inOrderSearch(t *testing.T) {
	tt := getTree()
	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	got := inOrderSearch(tt)
	assert.Equal(t, expected, got)
}
