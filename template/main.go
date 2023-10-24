package main

import "fmt"

var (
	smsOTP   *Sms
	emailOTP *Email
)

func Init() {
	smsOTP = &Sms{}
	emailOTP = &Email{}
}

func main() {
	Init()
	o := Otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)
	fmt.Println()
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
