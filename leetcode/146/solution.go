package main

import "fmt"

/*
146. LRU 缓存
请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

示例：

输入
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
输出
[null, null, null, 1, null, -1, null, -1, 3, 4]

解释
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // 缓存是 {1=1}
lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
lRUCache.get(1);    // 返回 1
lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
lRUCache.get(2);    // 返回 -1 (未找到)
lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
lRUCache.get(1);    // 返回 -1 (未找到)
lRUCache.get(3);    // 返回 3
lRUCache.get(4);    // 返回 4
*/
func main() {
	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(1, 1)
	lru.Put(2, 3)
	lru.Put(4, 1)
	fmt.Println(lru.cache)

	//getRet := lru.Get(1)
	//fmt.Println(getRet)

	//getRet := lru.Get(2)
	//fmt.Println(getRet)

	//lru.Put(4, 1)
	//fmt.Println(lru.cache)

	//getRet = lru.Get(1)
	//fmt.Println(getRet)
}

type LRUCache struct {
	capacity   int
	cache      map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	cache := make(map[int]*Node)
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    cache,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node := this.cache[key]
	if node == nil {
		return -1
	}
	this.afterAccess(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node := this.cache[key]
	if node != nil {
		node.val = value
		this.afterAccess(node)
		return
	}

	if len(this.cache) >= this.capacity {
		delKey := this.tail.prev.key

		this.tail.prev = this.tail.prev.prev
		this.tail.prev.next = this.tail

		delete(this.cache, delKey)
	}
	node = &Node{
		key:  key,
		val:  value,
		next: nil,
		prev: nil,
	}

	this.head.next.prev = node
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head

	this.cache[node.key] = node
}

func (this *LRUCache) afterAccess(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
	node.prev = this.head
}

type Node struct {
	key        int
	val        int
	next, prev *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("(%d, %d)", n.key, n.val)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
