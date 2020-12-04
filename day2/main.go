package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	err := os.Chdir("./")
	if err != nil {
		log.Fatalf("failed changing dir: %s", err)
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	pt1(txtlines)
	// pt2()

}
