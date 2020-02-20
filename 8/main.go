package main

import (
	"bufio"
	"fmt"
	"os"
)

var WIDTH = 25
var HEIGHT = 6

func main() {
	file, _ := os.Open("8/input.txt")
	defer file.Close()
	fileInfo, _ := file.Stat()
	scanner := bufio.NewScanner(file)

	// READ DATA
	i := 0
	data := make([]int, fileInfo.Size()-1)
	scanner.Scan() // only one line
	for _, char := range scanner.Text() {
		data[i] = int(char) - '0'
		i++
	}

	// SEPARATE LAYERS
	i = 0
	size := WIDTH * HEIGHT
	lim := len(data) / size
	layers := make([][]int, lim)
	layers[i] = make([]int, size)
	for j, char := range data {
		layers[i][j%size] = char
		if (j+1)%size == 0 {
			i++
			if i < lim {
				layers[i] = make([]int, size)
			}
		}
	}

	// PART A: IMAGE INTEGRITY
	var counts [3]int
	fewest := WIDTH * HEIGHT
	fewestProduct := 0
	for _, layer := range layers {
		counts = [3]int{0, 0, 0}
		for _, char := range layer {
			counts[char] += 1
		}
		if counts[0] < fewest {
			fewest = counts[0]
			fewestProduct = counts[1] * counts[2]
		}
	}

	fmt.Println(fewestProduct)

	// PART B: RENDER IMAGE
	image := layers[0]
	for _, layer := range layers {
		for i, pixel := range layer {
			if image[i] == 2 {
				image[i] = pixel
			}
		}
	}

	for i, pixel := range image {
		if pixel == 0 {
			fmt.Print(" ")
		} else {
			fmt.Print("█")
		}
		if (i+1)%WIDTH == 0 {
			fmt.Println()
		}
	}
}
