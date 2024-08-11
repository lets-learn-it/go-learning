package main

import "fmt"

type Student struct {
	Name  string
	Roll  int
	Marks int
}

func main() {
	integers()

	students()
}

func integers() {
	// array of integers
	numbers := []int{1, 2, 3, 4, 5}

	// find the index of 3
	index := Index(numbers, func(n int) bool {
		return n == 3
	})
	fmt.Println(index) // 2

	// check if all numbers are even
	allEven := All(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(allEven) // false

	// check if any number is odd
	anyOdd := Any(numbers, func(n int) bool {
		return n%2 != 0
	})
	fmt.Println(anyOdd) // true

	// filter even numbers
	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evenNumbers) // [2 4]

	// square each number
	squaredNumbers := Map(numbers, func(n int) int {
		return n * n
	})
	fmt.Println(squaredNumbers) // [1 4 9 16 25]
}

func students() {
	// array of students
	students := []Student{
		{"Alice", 1, 90},
		{"Bob", 2, 30},
		{"Charlie", 3, 55},
	}

	// find the index of student with roll number 2
	index := Index(students, func(s Student) bool {
		return s.Roll == 2
	})
	fmt.Println(index) // 1

	// check if all student passed
	allPassed := All(students, func(s Student) bool {
		return s.Marks >= 35
	})
	fmt.Println(allPassed) // false

	// check if any student failed
	anyFailed := Any(students, func(s Student) bool {
		return s.Roll < 35
	})
	fmt.Println(anyFailed) // true

	// filter passed students
	studentsPassed := Filter(students, func(s Student) bool {
		return s.Marks >= 35
	})
	fmt.Println(studentsPassed) // [{Alice 1 90} {Charlie 3 55}]

	// give 10 grace to all students
	updatedStudents := Map(students, func(s Student) Student {
		s.Marks += 10
		return s
	})
	fmt.Println(updatedStudents) // [{Alice 1 100} {Bob 2 40} {Charlie 3 65}]
}
