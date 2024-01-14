package main

type Bitset struct {
	size uint32
	set  []byte
}

func NewBitset(size uint32) *Bitset {
	size = (size + 7) / 8
	return &Bitset{
		size: size * 8,
		set:  make([]byte, size),
	}
}

func (b *Bitset) Set(n uint32) {
	if n >= b.size {
		return
	}

	b.set[n/8] |= (1 << (n % 8))
}

func (b *Bitset) UnSet(n uint32) {
	if n >= b.size {
		return
	}

	b.set[n/8] &= ^(1 << (n % 8))
}

func (b *Bitset) IsSet(n uint32) bool {
	if n >= b.size {
		return false
	}

	return b.set[n/8]&(1<<(n%8)) != 0
}
