package main

import "fmt"

type Patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Department interface {
	execute(*Patient)
	setNext(Department)
}

// Reception 具体处理者
type Reception struct {
	next Department
}

type Doctor struct {
	next Department
}

type Medical struct {
	next Department
}

type Cashier struct {
	next Department
}

func (r *Reception) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(patient)
		return
	}
	fmt.Println("Reception registering patient")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *Reception) setNext(department Department) {
	r.next = department
}

func (r *Doctor) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Doctor registration already done")
		r.next.execute(patient)
		return
	}
	fmt.Println("Doctor checking patient")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *Doctor) setNext(department Department) {
	r.next = department
}

func (r *Medical) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Medical registration already done")
		r.next.execute(patient)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	patient.registrationDone = true
	r.next.execute(patient)
}

func (r *Medical) setNext(department Department) {
	r.next = department
}

func (r *Cashier) execute(patient *Patient) {
	if patient.registrationDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (r *Cashier) setNext(department Department) {
	r.next = department
}
