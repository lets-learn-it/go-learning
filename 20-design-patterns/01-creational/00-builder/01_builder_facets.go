package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// job
	Company, Position string
	AnnualIncome      int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

// utility methods fro PersonBuilder
func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(pin string) *PersonAddressBuilder {
	pab.person.Postcode = pin
	return pab
}

func (pjb *PersonJobBuilder) At(company string) *PersonJobBuilder {
	pjb.person.Company = company
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earns(income int) *PersonJobBuilder {
	pjb.person.AnnualIncome = income
	return pjb
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	pb := NewPersonBuilder()

	person := pb.
		Lives().
		At("Sanjivani Poultry Farm").
		In("Kavathe Ekand").
		WithPostcode("416307").
		Works().
		At("Siemens").
		AsA("Software Engineer").
		Earns(0).
		Build()

	fmt.Print(person)
}
