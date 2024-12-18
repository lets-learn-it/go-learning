package main

import "fmt"

// This is like 2 different interfaces
// type AdderFloat32 interface { Add(float32) float32 }
// type AdderFloat64 interface { Add(float64) float64 }
type Adder[T float32 | float64] interface {
	Add(T) T
}

type Length struct {
	l    float32
	unit string
}

func (l Length) Add(x float32) float32 {
	return l.l + x
}

type AtomicDistance struct {
	d    float64
	unit string
}

func (d AtomicDistance) Add(x float64) float64 {
	return d.d + x
}

func main() {
	l := Length{l: 2.4345, unit: "meter"}
	d := AtomicDistance{d: 0.000234, unit: "nanometer"}

	fmt.Println(l.Add(4.56))
	fmt.Println(d.Add(0.00034))
}
