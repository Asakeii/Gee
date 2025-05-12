package gee

import "strings"

type node struct {
	path     string // 只有这个结点是一个路径的最后一个节点时，path才会有实际值
	part     string // 当前节点的部分 例如/hello/doc中的hello
	children []*node
	isWild   bool // false:模糊查询
}

func (n *node) matchChild(part string) *node {
	// 查找下一层的节点有没有要找的
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node) matchChildren(part string) []*node {
	// 得到下一层的所有节点
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node) insertChild(path string, parts []string, height int) {
	// 插入节点 注意：n的child才是正在插入的节点，而不是n本身，n是已有的节点
	if len(parts) == height {
		n.path = path
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{
			part:   part,
			isWild: part[0] == '*' || part[0] == ':',
		}
		n.children = append(n.children, child)
	}
	child.insertChild(path, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.path == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
