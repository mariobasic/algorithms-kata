package _lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU_detach(t *testing.T) {
	lru := newLRU[string, int](3)

	_, ok := lru.get("foo")
	assert.False(t, ok)

	lru.update("foo", 69)
	val, ok := lru.get("foo")
	assert.True(t, ok)
	assert.Equal(t, 69, val)

	lru.update("bar", 420)
	val, ok = lru.get("bar")
	assert.True(t, ok)
	assert.Equal(t, 420, val)

	lru.update("baz", 1337)
	val, ok = lru.get("baz")
	assert.True(t, ok)
	assert.Equal(t, 1337, val)

	lru.update("ball", 69420)

	val, ok = lru.get("ball")
	assert.True(t, ok)
	assert.Equal(t, 69420, val)

	val, ok = lru.get("foo")
	assert.False(t, ok)

	val, ok = lru.get("bar")
	assert.True(t, ok)
	assert.Equal(t, 420, val)

	lru.update("foo", 69)
	val, ok = lru.get("bar")
	assert.True(t, ok)
	assert.Equal(t, 420, val)

	val, ok = lru.get("foo")
	assert.True(t, ok)
	assert.Equal(t, 69, val)

	val, ok = lru.get("baz")
	assert.False(t, ok)
}
