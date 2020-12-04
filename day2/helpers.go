package main

import (
	"fmt"
	"strconv"
	"strings"
)

func (p PWData) pwValidator() bool {
	count := strings.Count(p.Password, p.Letter)

	if count < p.MinCount || count > p.MaxCount {
		return false
	}
	return true
}

// PWDataBuilder takes the string and constructs the struct type to be returned
func PWDataBuilder(input string) PWData {
	xs := strings.Split(input, ":")
	rules := xs[0]
	pw := xs[1]

	vals := strings.Split(rules, " ")

	letter := vals[1]
	fmt.Println("letter", letter)
	minmax := strings.Split(vals[0], "-")
	fmt.Println("min max", minmax)
	maxcount, _ := strconv.Atoi(minmax[1])
	mincount, _ := strconv.Atoi(minmax[0])

	pwd := PWData{
		Letter:   letter,
		MaxCount: maxcount,
		MinCount: mincount,
		Password: pw,
	}

	return pwd
}
