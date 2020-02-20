package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("1/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		total += mass/3 - 2
	}
	fmt.Println(total)

	file.Seek(0, 0)
	scanner = bufio.NewScanner(file)
	total = 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		for mass > 0 {
			mass = mass/3 - 2
			if mass < 0 {
				mass = 0
			}
			total += mass
		}
	}
	fmt.Println(total)
}
