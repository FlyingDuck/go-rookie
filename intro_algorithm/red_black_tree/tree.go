package red_black_tree

type RedBlackTree struct {
	root *RBTreeNode
}

// LeftRotate 左旋
// x：要旋转的节点
func (T *RedBlackTree) LeftRotate(x *RBTreeNode) {
	y := x.right
	// 确定中间节点位置
	x.right = y.left
	if !y.left.isNil() {
		y.left.parent = x
	}
	// 确定 y 的父节点（链接 x 的父节与 y）
	y.parent = x.parent
	if x.parent.isRoot() { // x 是根节点
		T.root = y
	} else if x.isLeft() { // x 是左节点
		x.parent.left = y
	} else { // x 是右节点
		x.parent.right = y
	}
	// 确定 y 的左节点
	y.left = x
	x.parent = y
}

// RightRotate 右旋
// y：要旋转的节点
func (T *RedBlackTree) RightRotate(y *RBTreeNode) {
	x := y.left
	//
	y.left = x.right
	if !x.right.isNil() {
		x.right.parent = y
	}
	//
	x.parent = y.parent
	if y.parent.isRoot() {
		T.root = x
	} else if y.isLeft() {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	//
	x.right = y
	y.parent = x
}
