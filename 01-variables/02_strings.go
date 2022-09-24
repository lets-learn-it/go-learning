package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

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
}
