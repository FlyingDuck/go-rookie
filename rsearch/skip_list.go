package rsearch

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	MAX_LEVEL = 16 //最高层数
)

//跳表节点结构体
type SkipListNode struct {
	v        interface{}     //跳表保存的值
	score    int             //用于排序的分值
	level    int             //层高
	forwards []*SkipListNode //每层前进指针
}

//新建跳表节点
func newSkipListNode(v interface{}, score, level int) *SkipListNode {
	return &SkipListNode{
		v:        v,
		score:    score,
		forwards: make([]*SkipListNode, level, level),
		level:    level,
	}
}

//跳表结构体
type SkipList struct {
	head   *SkipListNode //跳表头结点
	level  int           //跳表当前层数
	length int           //跳表长度
}

//实例化跳表对象
func NewSkipList() *SkipList {
	//头结点，便于操作
	head := newSkipListNode(0, math.MinInt32, MAX_LEVEL)
	return &SkipList{head, 1, 0}
}

//获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

//获取跳表层级
func (sl *SkipList) Level() int {
	return sl.level
}

//查找
func (sl *SkipList) Find(v interface{}, score int) *SkipListNode {
	if nil == v || sl.length == 0 {
		return nil
	}

	cur := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

//插入节点到跳表中
func (sl *SkipList) Insert(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找插入位置
	cur := sl.head
	//记录每层的路径
	update := [MAX_LEVEL]*SkipListNode{}
	i := MAX_LEVEL - 1
	for ; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if nil == cur.forwards[i] {
			update[i] = cur
		}
	}

	//通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	//创建一个新的跳表节点
	newNode := newSkipListNode(v, score, level)

	//原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

//删除节点
func (sl *SkipList) Delete(v interface{}, score int) int {
	if nil == v {
		return 1
	}

	//查找前驱节点
	cur := sl.head
	//记录要删除节点的所有前驱节点
	update := [MAX_LEVEL]*SkipListNode{}
	for i := sl.level - 1; i >= 0; i-- {
		update[i] = sl.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		// 如果删除的节点的前驱节点是head，并且没有后继节点，说明是顶层唯一节点，需要改变跳表的高度
		if update[i] == sl.head && cur.forwards[i] == nil {
			sl.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	sl.length--

	return 0
}

func (sl *SkipList) String() string {
	return fmt.Sprintf("level:%d, length:%d", sl.level, sl.length)
}
