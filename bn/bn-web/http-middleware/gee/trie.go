package gee

import (
	"fmt"
	"strings"
)

type node struct {
	pattern  string  // 待匹配的路由
	part     string  // 路由中的一部分
	children []*node // 子节点
	isFuzzy  bool    // 是否是模糊匹配
}

func parserPath(path string) []string {
	if path == "" {
		return []string{}
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return strings.Split(strings.TrimSuffix(path, "/"), "/")[1:]
}

func (n *node) insert(pattern string, parts []string, level int) {
	if len(parts) == level {
		n.pattern = pattern
		return
	}
	part := parts[level]
	childNode := n.matchInsertChild(part)
	if childNode == nil {
		childNode = &node{
			part:    part,
			isFuzzy: strings.HasPrefix(part, ":") || strings.HasPrefix(part, "*"),
		}
		n.children = append(n.children, childNode)
	}
	childNode.insert(pattern, parts, level+1)
}

func (n *node) matchInsertChild(part string) *node {
	for _, childNode := range n.children {
		if childNode.part == part {
			return childNode
		}
	}
	return nil
}

func (n *node) search(reqParts []string, level int) *node {
	fmt.Printf("level : %d, %s\n", level, reqParts)
	if len(reqParts) == level || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	children := n.matchSearchChild(reqParts[level])
	for _, childNode := range children {
		return childNode.search(reqParts, level+1)
	}
	return nil

}

func (n *node) matchSearchChild(reqPart string) []*node {
	nodes := make([]*node, 0)
	for _, childNode := range n.children {
		if childNode.isFuzzy || childNode.part == reqPart {
			nodes = append(nodes, childNode)
		}
	}
	return nodes
}

func (n *node) printAllChildString() {
	fmt.Printf("node{path = %s , part=%s, isFuzzy=%t}\n", n.pattern, n.part, n.isFuzzy)
	if n.children == nil || len(n.children) < 1 {
		return
	}
	for _, childNode := range n.children {
		childNode.printAllChildString()
	}
}
