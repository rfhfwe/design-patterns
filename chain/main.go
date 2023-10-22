package main

var (
	casher    *Cashier
	medcial   *Medical
	doctor    *Doctor
	reception *Reception
	patient   *Patient
)

func Init() {
	casher = &Cashier{}

	medcial = &Medical{}
	medcial.setNext(casher)

	doctor = &Doctor{}
	doctor.setNext(medcial)

	reception = &Reception{}
	reception.setNext(doctor)

	patient = &Patient{name: "abc"}
}

func main() {
	Init() // 初始化
	reception.execute(patient)
}
