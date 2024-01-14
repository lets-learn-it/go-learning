package main

import (
	"flag"
	"fmt"

	"github.com/pkg/profile"
)

var s = flag.Int("n", 2000, "help message for flag n")

func main() {
	defer profile.Start(profile.MemProfileRate(512), profile.ProfilePath(".")).Stop()
	var i uint32 = 0
	size := uint32(*s)
	b := NewBitset(size)

	for i < size {
		b.Set(i)
		i += 2
	}

	i = 0
	for i < size {
		if !b.IsSet(i) {
			fmt.Printf("pos %d is not set\n", i)
		}

		b.UnSet(i)

		if b.IsSet(i) {
			fmt.Printf("pos %d is set\n", i)
		}

		i += 2
	}
}
