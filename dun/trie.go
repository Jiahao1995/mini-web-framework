package dun

import (
	"fmt"
	"strings"
)

type Node struct {
	pattern  string
	part     string
	children []*Node
	isWild   bool
}

func (node *Node) string() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", node.pattern, node.part, node.isWild)
}

func (node *Node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		node.pattern = pattern
		return
	}
	part := parts[height]
	child := node.matchChild(part)
	if child == nil {
		child = &Node{
			part:   part,
			isWild: part[0] == ':' || part[0] == '*',
		}
		node.children = append(node.children, child)
	}
	child.insert(pattern, parts, height+1)
}

func (node *Node) search(parts []string, height int) *Node {
	if len(parts) == height || strings.HasPrefix(node.part, "*") {
		if node.pattern == "" {
			return nil
		}
		return node
	}
	part := parts[height]
	children := node.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (node *Node) travel(list *[]*Node) {
	if node.pattern != "" {
		*list = append(*list, node)
	}
	for _, child := range node.children {
		child.travel(list)
	}
}

func (node *Node) matchChild(part string) *Node {
	for _, child := range node.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (node *Node) matchChildren(part string) []*Node {
	nodes := make([]*Node, 0)
	for _, child := range node.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
