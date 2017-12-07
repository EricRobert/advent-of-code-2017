package p7

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func Main(args []string) {
	if len(args) != 2 {
		log.Fatal("usage: advent-of-code-1027 7[a|b] 'filename'")
	}

	switch args[0] {
	case "7a", "7":
		fmt.Print(Root(args[1]))
	}
}

type Nodes map[string]*Node

type Node struct {
	Name     string
	Parent   string
	Children []string
}

func (m Nodes) Add(line string) {
	p := strings.Split(line, " ")

	get := func(name string) *Node {
		node, ok := m[name]
		if !ok {
			node = &Node{Name: name}
			m[name] = node
		}

		return node
	}

	node := get(p[0])

	if len(p) == 2 || len(p) == 3 {
		return
	}

	for _, i := range p[3:] {
		i = strings.Trim(i, ",")
		node.Children = append(node.Children, i)
		child := get(i)
		if child.Parent != "" {
			log.Fatal("more than 1 parent")
		}

		child.Parent = node.Name
	}
}

func Root(filename string) string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(f), "\n")
	nodes := Nodes{}
	for _, line := range lines {
		if line == "" {
			continue
		}

		nodes.Add(line)
	}

	root := ""
	for k, v := range nodes {
		if v.Parent == "" {
			if root != "" {
				log.Fatal("more than 1 root")
			}

			root = k
		}
	}

	return root
}
