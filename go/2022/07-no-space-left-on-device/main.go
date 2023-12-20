package nospaceleftondevice07

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	c "github.com/patricho/advent-of-code/go/util"
)

type Node struct {
	name     string
	level    int
	size     int
	isDir    bool
	children []*Node
	parent   *Node
}

var output = c.ReadFile("07-no-space-left-on-device/input.txt")

func Part2() {
	root := Part1()

	totalSpace := 70_000_000
	freeSpace := totalSpace - root.size
	neededSpace := 30_000_000 - freeSpace

	fmt.Println("root.size", root.size, "freeSpace", freeSpace, "neededSpace", neededSpace)

	folders := findLargeFolders(root, neededSpace)

	sort.Sort(BySize(folders))
	for _, f := range folders {
		fmt.Println("larger than", neededSpace, f.name, f.size)
	}
}

func findLargeFolders(node *Node, neededSpace int) []*Node {
	folders := []*Node{}

	for _, ch := range node.children {
		if ch.isDir {
			if ch.size >= neededSpace {
				folders = append(folders, ch)
			}

			folders = append(folders, findLargeFolders(ch, neededSpace)...)
		}
	}

	return folders
}

func Part1() *Node {
	root := Node{
		name:     "/",
		level:    0,
		size:     0,
		isDir:    true,
		children: []*Node{},
	}

	process(&root, 0)

	countSizes(&root)

	//traverse(&root)
	//
	//answer := findSmallFolders(&root)
	//
	//fmt.Println("answer", answer)

	return &root
}

func findSmallFolders(node *Node) int {
	sum := 0

	for _, ch := range node.children {
		if ch.isDir {
			if ch.size <= 100_000 {
				fmt.Println("small enough", ch.name, ch.size)
				sum += ch.size
			}

			sum += findSmallFolders(ch)
		}
	}

	return sum
}

func countSizes(node *Node) int {
	sum := 0

	for _, ch := range node.children {
		if ch.isDir {
			sum += countSizes(ch)
		} else {
			sum += ch.size
		}
	}

	node.size = sum

	return sum
}

func process(node *Node, idx int) {
	if idx >= len(output) {
		return
	}

	// fmt.Println("processing", node.name, idx, output[idx])

	// process current command
	if output[idx] == "$ cd /" {
		// just getting started yo
		idx++
	} else if strings.HasPrefix(output[idx], "$ cd") {
		re, _ := regexp.Compile(`^\$ cd (.+)$`)
		dir := re.ReplaceAllString(output[idx], "$1")

		if dir == ".." {
			process(node.parent, idx+1)
			return
		}

		newnode := findNodeByName(node, dir)
		process(newnode, idx+1)
		return
	} else if output[idx] == "$ ls" {
		// listing files

		for {
			idx++
			if idx >= len(output) || output[idx][0] == '$' {
				// reached the end of file listing
				break
			}

			re, _ := regexp.Compile(`^(dir|\d+) (.+)$`)

			dirOrSize := re.ReplaceAllString(output[idx], "$1")
			name := re.ReplaceAllString(output[idx], "$2")

			if dirOrSize == "dir" {
				node.children = append(node.children, &Node{
					name:     name,
					level:    node.level + 1,
					size:     0,
					isDir:    true,
					children: []*Node{},
					parent:   node,
				})
			} else {
				size, _ := strconv.Atoi(dirOrSize)
				node.children = append(node.children, &Node{
					name:   name,
					level:  node.level + 1,
					size:   size,
					isDir:  false,
					parent: node,
				})
			}
		}
	} else {
		// temp
		idx++
	}

	// proceed
	process(node, idx)
}

func findNodeByName(node *Node, dir string) *Node {
	for _, n := range node.children {
		if n.isDir && n.name == dir {
			return n
		}
	}

	return nil
}

func traverse(node *Node) {
	fmt.Println(strings.Repeat(" ", node.level*4), node.name, node.size)

	if node.name == "b.txt" {
		node.size = 666
	}

	if node.children == nil {
		return
	}

	for _, c := range node.children {
		traverse(c)
	}
}

type BySize []*Node

func (a BySize) Len() int           { return len(a) }
func (a BySize) Less(i, j int) bool { return a[i].size < a[j].size }
func (a BySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
