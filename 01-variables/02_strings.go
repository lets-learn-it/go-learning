package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/*
 * byte: a synonym for uint8
 * rune: a synonym for int32 for character
 * string: immutable sequence of characters
 *    - physically a sequence of bytes (UTF-8 encoding)
 *    - logically a sequence of unicode runes
 * strings are immutable and can share undelying storage
 */

func main() {
	// string literal
	name := "Parikshit"
	path := "c:\\user\\parikshit"

	// raw string literal
	myName := `Parikshit
						patil`
	myPath := `c:\user\parikshit`

	fmt.Printf("%s\n%s\n", name, myName)

	fmt.Printf("%s\n%s\n", path, myPath)

	// len actually prints bytes not length
	myNameMarathi := "परिक्षित"
	fmt.Println(len(myNameMarathi))
	fmt.Println(utf8.RuneCountInString(myNameMarathi))

	lotOfSpaces := `
	This sentences

	  has lot of spaces    `
	fmt.Println("Trimmed version (only leading & training): " + strings.TrimSpace(lotOfSpaces))

	fmt.Printf("Type: %8T %[1]v\n", name)
	fmt.Printf("Type: %8T %[1]v\n", myNameMarathi)
	fmt.Printf("%8T %[1]v\n", []rune(myNameMarathi))
	fmt.Printf("%8T %[1]v\n", []byte(myNameMarathi))

	// substrings
	fmt.Println(name[3:] + name[:3])

	// helper functions
	fmt.Println(strings.Contains(name, "ari"))
	fmt.Println(strings.Contains(name, "aris"))
	fmt.Println(strings.HasPrefix(name, "Pari"))
	fmt.Println(strings.Index(name, "ari"))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.Split(path, "\\"))
	fmt.Println(strings.Join(strings.Split(path, "\\"), "/"))
	fmt.Println(len(strings.Split(" hello there   whats up  ", " ")))
}
