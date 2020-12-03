package main

import "fmt"

func pt2(nums []int) {
	// data := make(map[int]struct{})
	diffs := make(map[int]struct{})

	for outter := range nums {
		for inner := range nums {
			if outter != inner {
				diffs[final-outter] = struct{}{}
			}
		}
	}

	for d := range diffs {
		fmt.Println("d", d)
	}

	// fmt.Println(diffs)

}
