package main

import (
	"../utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEBUG = true

/*
9 ORE => 2 A
8 ORE => 3 B
7 ORE => 5 C
3 A, 4 B => 1 AB
5 B, 7 C => 1 BC
4 C, 1 A => 1 CA
2 AB, 3 BC, 4 CA => 1 FUEL

             1FUEL
            /  |  \
         2AB  3BC  4CA
        / |   / \   | \
      3A 4B  5B 7C  4C 1A = 6A+8B+15B+21C+16C+4A = 10A + 23B + 37C
      |   |  |   |  |   |
      9   8  8   7  9   8 = 10/2*9 + 23/3*8 + 37/5*7 = 45 + 64 + 56 = 165
*/

type tree struct{ utils.Tree }

type data struct {
	quantity int
	costs    []int
}

func (t *tree) addNodeFromLine(s string) {
	l := strings.Split(s, " => ")
	product, requirements := l[1], strings.Split(l[0], ", ")

	l = strings.Split(product, " ")
	key := l[1]
	quantity, _ := strconv.Atoi(l[0])
	n := t.Get(key)
	d := data{}
	d.quantity = quantity
	d.costs = make([]int, len(requirements))

	n.Children = make([]*utils.Node, len(requirements))
	for i, s := range requirements {
		l := strings.Split(s, " ")
		n.Children[i] = t.Get(l[1])
		n.Children[i].Parent = n
		d.costs[i], _ = strconv.Atoi(l[0])
	}

	n.Data = d
}

func newTree(inputfile string) *tree {
	t := tree{*utils.NewTree()}

	file, _ := os.Open(inputfile)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Build tree
	for scanner.Scan() {
		t.addNodeFromLine(scanner.Text())
	}

	return t
}

func (t *tree) getCost() {
	costs := make(map[string]int)
}

func main() {
	if DEBUG {
		t := newTree("14/test0.txt")

	}

	fmt.Println(0)
}
