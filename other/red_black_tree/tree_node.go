package red_black_tree

type Color int32

const (
	BLACK Color = 0
	RED   Color = 1
)

var (
	Nil = RBTreeNode{
		color: BLACK,
	}
)

type RBTreeNode struct {
	color  Color
	key    interface{}
	left   *RBTreeNode
	right  *RBTreeNode
	parent *RBTreeNode
}

func (node *RBTreeNode) isNil() bool {
	return node == &Nil
}

func (node *RBTreeNode) isRoot() bool {
	return !node.isNil() && node.parent.isNil()
}

func (node *RBTreeNode) isLeft() bool {
	return !node.isRoot() && node == node.parent.left
}

func (node *RBTreeNode) isRight() bool {
	return !node.isRoot() && node == node.parent.right
}
