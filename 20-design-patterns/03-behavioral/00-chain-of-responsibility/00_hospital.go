package main

import "fmt"

type Department interface {
	handle(p *Patient)
	setNext(d Department)
}

type Patient struct {
	name string
}

// abstract department which will implement setNext method
type DefaultDepartment struct {
	next Department
}

func (d *DefaultDepartment) setNext(next Department) {
	d.next = next
}

// departments
type Reception struct {
	DefaultDepartment
}

func (r *Reception) handle(p *Patient) {
	fmt.Printf("Patient %s is received by Reception\n", p.name)
	if r.next != nil {
		r.next.handle(p)
	}
}

type Doctor struct {
	DefaultDepartment
}

func (d *Doctor) handle(p *Patient) {
	fmt.Printf("Patient %s is examined by Doctor\n", p.name)
	if d.next != nil {
		d.next.handle(p)
	}
}

type Cashier struct {
	DefaultDepartment
}

func (c *Cashier) handle(p *Patient) {
	fmt.Printf("Patient %s is billed by Cashier\n", p.name)
	if c.next != nil {
		c.next.handle(p)
	}
}

func main() {
	reception := &Reception{}
	doctor := &Doctor{}
	cashier := &Cashier{}

	reception.setNext(doctor)
	doctor.setNext(cashier)

	patient := &Patient{name: "John Doe"}

	reception.handle(patient)
}
