package main

import (
	"testing"
)

func TestBitSet(t *testing.T) {
	var size uint32 = 43
	var i uint32 = 0
	b := NewBitset(size)

	for i < size {
		b.Set(i)
		i += 2
	}

	i = 0
	for i < size {
		if !b.IsSet(i) {
			t.Errorf("pos %d is not set\n", i)
		}

		b.UnSet(i)

		if b.IsSet(i) {
			t.Errorf("pos %d is set\n", i)
		}

		i += 2
	}
}
