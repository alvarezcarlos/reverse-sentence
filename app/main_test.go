package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseSentence(t *testing.T) {
	t.Run("When sentence is giving expect string reverse sentence word by word", func(t *testing.T) {
		sentence := "today is the last day of the rest of my life"
		r := reverseSentence(sentence)
		assert.Equal(t, "life my of rest the of day last the is today", r)
	})
}
