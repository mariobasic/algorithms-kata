package regular

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_postOrderSearch(t *testing.T) {
	tt := getTree()
	expected := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	got := postOrderSearch(tt)
	assert.Equal(t, expected, got)
}
