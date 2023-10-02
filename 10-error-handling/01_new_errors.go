package main

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

var ErrDivideByFour = errors.New("You are dividing by 4 which is not allowed")

type ErrFileOpen struct {
	filename string
	count    int8
}

func (e *ErrFileOpen) Error() string {
	return fmt.Sprintf("file: %s, count: %d", e.filename, e.count)
}

func Divide(a, b int) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("can't divide %d by 0", a)
	}

	if b == 4 {
		return 0, ErrDivideByFour
	}
	return float32(a) / float32(b), nil
}

func Open(filename string) (bool, error) {
	if utf8.RuneCountInString(filename) == 4 {
		return false, &ErrFileOpen{
			filename: filename,
			count:    int8(utf8.RuneCountInString(filename)),
		}
	}
	return true, nil
}

func main() {
	val, _ := Divide(43, 7)
	fmt.Println("43 / 7 = ", val)

	val, err := Divide(23, 0)
	fmt.Println(err)

	val, err = Divide(26, 4)
	fmt.Println(err)

	// check if returned error is ErrDivideByFour
	if errors.Is(err, ErrDivideByFour) {
		fmt.Println("Yes, it is divide by four...")
	}

	// open magical file
	_, err = Open("PSKP")
	var targetErr *ErrFileOpen
	if errors.As(err, &targetErr) {
		fmt.Println("Yeah, its ErrFileOpen", targetErr.filename)
	}
}
