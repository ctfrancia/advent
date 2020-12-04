package main

import (
	"fmt"
)

// PWData represents the data about the passwords
type PWData struct {
	Letter   string
	MinCount int
	MaxCount int
	Password string
}

func pt1(txt []string) {
	correctNum := 0
	xpwdata := make([]PWData, len(txt))

	for _, eachline := range txt {
		pwdchunk := PWDataBuilder(eachline)
		xpwdata = append(xpwdata, pwdchunk)
	}

	fmt.Println("after ->", xpwdata)
	for _, pwdata := range xpwdata {
		valid := pwdata.pwValidator()

		if valid {
			correctNum++
		}
	}
	fmt.Println(correctNum)
}
