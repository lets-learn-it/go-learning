package main

import "fmt"

func main() {

	/* Wrong
	 * because i is single memory location & all functions are refering to that single i
	 * thats why at last i is holding 4 value. All function will print 4 & address of same i
	 */
	s1 := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s1[i] = func() {
			fmt.Printf("%d @ %p\n", i, &i)
		}
	}

	for i := 0; i < 4; i++ {
		s1[i]()
	}

	/* Right
	 * Create another variable i2 for avoiding above situation.
	 * there will be new i2 created for each function call
	 */
	s2 := make([]func(), 4)

	for i := 0; i < 4; i++ {
		i2 := i // closure capture
		s2[i] = func() {
			fmt.Printf("%d @ %p\n", i2, &i2)
		}
	}

	for i := 0; i < 4; i++ {
		s2[i]()
	}
}
