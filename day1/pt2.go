package main

import "fmt"

func pt1(nums []int) {
	s := make(map[int]int)

	for _, num := range nums {
		dif := final - num

		if s[dif] == dif {
			fmt.Println("found", s[dif], num)
			solution := s[dif] * num
			fmt.Println("solution", solution)
		}
		s[num] = num
	}
}
