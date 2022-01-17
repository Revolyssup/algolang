package tree

type NodeCallbackFn func(n *Node) error

type Tree interface {
	Traverse(NodeCallbackFn, int) error
	Type() string
	Join(Tree) error
}
