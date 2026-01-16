package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
  assert.Equal(t, 15, Add(8, 7))
  assert.Equal(t, 1, Add(1, 0))
  assert.Equal(t, 10, Add(9, 1))
}
