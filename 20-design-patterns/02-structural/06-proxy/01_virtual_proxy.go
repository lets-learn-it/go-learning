package main

import "fmt"

// original
type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(f string) *Bitmap {
	fmt.Printf("loading bitmap image %s\n", f)
	return &Bitmap{filename: f}
}

func (b *Bitmap) Draw() {
	fmt.Printf("drawing bitmap %s\n", b.filename)
}

// proxy
type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(f string) *LazyBitmap {
	return &LazyBitmap{filename: f}
}

func (b *LazyBitmap) Draw() {
	if b.bitmap == nil {
		b.bitmap = NewBitmap(b.filename)
	}

	b.bitmap.Draw()
}

func main() {
	// Without lazy loading
	_ = NewBitmap("somefile.png")
	// bitmap.Draw()

	// lazyload
	_ = NewLazyBitmap("somefile.png")
}
