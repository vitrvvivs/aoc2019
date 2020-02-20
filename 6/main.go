package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type orbitNode struct {
	name     string
	depth    int
	parent   *orbitNode
	children []*orbitNode
}

func NeworbitNode(name string) *orbitNode {
	this := orbitNode{}
	this.name = name
	this.children = make([]*orbitNode, 0)
	return &this
}
func (this *orbitNode) addParent(other *orbitNode) {
	this.parent = other
	other.children = append(other.children, this)
}
func (this *orbitNode) walk(f func(*orbitNode)) {
	f(this)
	for _, child := range this.children {
		child.walk(f)
	}
}
func (this *orbitNode) SetDepth(depth int) {
	this.depth = depth
	for _, child := range this.children {
		child.SetDepth(depth + 1)
	}
}
func (this *orbitNode) SetDistance(distance int, ignore *orbitNode) {
	// same as depth, just include parent in "children"
	this.depth = distance
	for _, node := range append(this.children, this.parent) {
		if node != nil && node != ignore {
			node.SetDistance(distance+1, this)
		}
	}
}

func AddPair(this map[string]*orbitNode, parentName string, childName string) {
	parent, parentExists := this[parentName]
	child, childExists := this[childName]
	if !parentExists {
		parent = NeworbitNode(parentName)
		this[parentName] = parent
	}
	if !childExists {
		child = NeworbitNode(childName)
		this[childName] = child
	}

	child.addParent(parent)
}

func main() {
	file, _ := os.Open("6/input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Build tree
	index := make(map[string]*orbitNode)
	for scanner.Scan() {
		l := strings.Split(scanner.Text(), ")")
		AddPair(index, l[0], l[1])
	}

	// A
	sum := 0
	root := index["COM"]
	root.SetDepth(0)
	addDepths := func(node *orbitNode) { sum += node.depth }
	root.walk(addDepths)
	fmt.Println(sum)

	// B
	index["YOU"].SetDistance(0, nil)
	fmt.Println(index["SAN"].depth - 2)
}
