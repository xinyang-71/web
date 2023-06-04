package frameWork

import "strings"

// 用来支持对路由树的操作
type router struct {
	//GET有一棵树 POST有一棵树

	//http method =>路由树根节点
	trees map[string]*node
}

type node struct {
	path string

	// 子path到子节点
	children map[string]*node

	//用户注册的业务逻辑
	handler HandleFunc
}

// 第一个是正确的子节点，第二个是
//func (n *node) childOf(seg string) (*node, bool) {
//
//}

func (n *node) childOrCreate(seg string) *node {
	if n.children == nil {
		n.children = map[string]*node{}
	}
	res, ok := n.children[seg]
	if !ok {
		res = &node{
			path: seg,
		}
		n.children[seg] = res
	}
	return res
}

func (r *router) AddRoute(method string, path string, handleFunc HandleFunc) {
	// 首先要找到树
	root, ok := r.trees[method]
	if !ok {
		// 还没根节点
		root = &node{
			path: "/",
		}
		r.trees[method] = root
	}
	//path = path[1:]
	// 切割这个path
	segs := strings.Split(path, "/")
	for _, seg := range segs {
		// 递归下去找准位置，中途有节点不存在就要创建
		children := root.childOrCreate(seg)
		root = children
	}
	root.handler = handleFunc
}

func newRouter() *router {
	return &router{
		trees: map[string]*node{},
	}
}
