package main

import "fmt"

type Bits uint32
// https://yourbasic.org/golang/bitmask-flag-set-clear/

func main(){

	const (
		First Bits = 1 << iota  //1
		Second  //2
		Third   //4
		Fourth  //8
		Fifth //16
		Sixth  //32
		Seventh //64
		Eighth //128
		Ninth //256
	)

	var status Bits
	//status |= ParentAuction
	status |= First
	status |= Second
	fmt.Println("status First, Second: ", status) //6

	fmt.Printf("status has First: %t\n",Has(status, First))
	fmt.Printf("status has Second: %t\n",Has(status, Second))
	fmt.Printf("status has Third: %t\n",Has(status, Third))
}

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }