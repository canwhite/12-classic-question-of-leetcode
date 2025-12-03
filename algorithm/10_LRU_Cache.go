package algorithm

// 双向链表节点
type Node struct {
    key   int
    value int
    prev  *Node
    next  *Node
}

// 双向链表结构
type DoubleLinkedList struct {
    head *Node
    tail *Node
    size int
}

// LRU缓存结构
type LRUCache struct {
    capacity     int               // 缓存容量
    cache        map[int]*Node     // 哈希表：key -> 节点
    linkedList   *DoubleLinkedList // 双向链表：维护访问顺序
}

// 创建LRU缓存
func Constructor(capacity int) LRUCache {
    return LRUCache{
        capacity:   capacity,
        cache:      make(map[int]*Node),
        linkedList: &DoubleLinkedList{},
    }
}

// 从缓存中获取值
func (this *LRUCache) Get(key int) int {
    // 如果key不存在，返回-1
    if node, exists := this.cache[key]; exists {
        // 关键：访问后要移到链表头部（表示最近使用）
        this.linkedList.moveToHead(node)
        return node.value
    }
    return -1
}

// 向缓存中放入键值对
func (this *LRUCache) Put(key int, value int) {
    // 如果key已存在，更新值并移到头部
    if node, exists := this.cache[key]; exists {
        node.value = value
        this.linkedList.moveToHead(node)
        return
    }

    // 如果key不存在且缓存已满，删除最久未使用的节点
    if len(this.cache) >= this.capacity {
        // 删除链表尾部节点（最久未使用）
        tailNode := this.linkedList.removeTail()
        delete(this.cache, tailNode.key)
    }

    // 创建新节点并添加到头部
    newNode := &Node{key: key, value: value}
    this.linkedList.addToHead(newNode)
    this.cache[key] = newNode
}

// 双向链表的方法

// 添加节点到头部
func (dll *DoubleLinkedList) addToHead(node *Node) {
    node.prev = nil
    node.next = dll.head

    if dll.head != nil {
        dll.head.prev = node
    } else {
        // 链表为空，新节点也是尾部
        dll.tail = node
    }

    dll.head = node
    dll.size++
}

// 移除指定节点
func (dll *DoubleLinkedList) removeNode(node *Node) {
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        // 移除的是头节点
        dll.head = node.next
    }

    if node.next != nil {
        node.next.prev = node.prev
    } else {
        // 移除的是尾节点
        dll.tail = node.prev
    }

    dll.size--
}

// 将节点移动到头部
func (dll *DoubleLinkedList) moveToHead(node *Node) {
    dll.removeNode(node)
    dll.addToHead(node)
}

// 移除尾部节点
func (dll *DoubleLinkedList) removeTail() *Node {
    if dll.tail == nil {
        return nil
    }

    tailNode := dll.tail
    dll.removeNode(tailNode)
    return tailNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */