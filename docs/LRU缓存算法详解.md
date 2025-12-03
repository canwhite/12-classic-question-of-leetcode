# LRU缓存算法详解

## 🎯 LRU是什么？

### **LRU (Least Recently Used)** - 最近最少使用

**核心思想**：当缓存满了，优先删除**最久没有被访问**的数据。

**生活中的例子**：
```
想象你的书桌只能放5本书：
1. 你最近在看《算法导论》→ 放在桌面最显眼位置
2. 昨天看了《数据结构》→ 放在第二层
3. 上周看了《操作系统》→ 放在第三层
...
当你第6次拿新书来，会把最久没碰的那本放回书架

这就是LRU！
```

### **LRU的三大特点**

1. **时效性**：最近使用的数据更有价值
2. **有限容量**：缓存有大小限制，需要淘汰策略
3. **快速访问**：O(1)时间复杂度的get和put操作

## 🔧 技术实现

### **两个核心数据结构**

1. **哈希表**：`map[int]*Node`
   - 作用：O(1)时间查找 key → 节点
   - 存储键值对的映射关系

2. **双向链表**：`head ←→ ←→ ←→ tail`
   - 作用：维护访问顺序
   - `head`：最近使用的数据
   - `tail`：最久未使用的数据

### **为什么选择双向链表？**

**双向链表的关键优势**：
- **O(1)删除任意节点**：通过节点的prev指针直接删除
- **O(1)移动到头部**：断开节点，重新插入头部

如果用单向链表：
- 删除节点需要O(n)时间（从头部遍历找到前一个节点）

## 🔄 详细工作流程

### **工作流程详解**

```
初始状态（容量=2）：
缓存：{}
链表：head ←→ tail

1. Put(1, "A")
   缓存：{1: "A"}
   链表：head ←→ [1:"A"] ←→ tail

2. Put(2, "B")
   缓存：{1: "A", 2: "B"}
   链表：head ←→ [2:"B"] ←→ [1:"A"] ←→ tail

3. Get(1) ← 访问key=1
   步骤：
   - 哈希表找到1 → 节点[1:"A"]
   - 将节点移动到头部
   缓存：{1: "A", 2: "B"}
   链表：head ←→ [1:"A"] ←→ [2:"B"] ←→ tail

4. Put(3, "C") ← 缓存已满！
   步骤：
   - 删除尾部节点[2:"B"]（最久未使用）
   - 添加新节点[3:"C"]到头部
   缓存：{1: "A", 3: "C"}
   链表：head ←→ [3:"C"] ←→ [1:"A"] ←→ tail
```

### **Get操作详细步骤**

```go
func (this *LRUCache) Get(key int) int {
    // 步骤1：检查key是否存在
    if node, exists := this.cache[key]; exists {
        // 步骤2：存在，将节点移到头部（标记为最近使用）
        this.linkedList.moveToHead(node)
        return node.value
    }
    // 步骤3：不存在，返回-1
    return -1
}
```

### **Put操作详细步骤**

```go
func (this *LRUCache) Put(key int, value int) {
    // 情况1：key已存在
    if node, exists := this.cache[key]; exists {
        node.value = value
        this.linkedList.moveToHead(node)
        return
    }

    // 情况2：key不存在，但缓存已满
    if len(this.cache) >= this.capacity {
        // 删除最久未使用的节点（尾部）
        tailNode := this.linkedList.removeTail()
        delete(this.cache, tailNode.key)
    }

    // 情况3：添加新节点
    newNode := &Node{key: key, value: value}
    this.linkedList.addToHead(newNode)
    this.cache[key] = newNode
}
```

## 🏗️ 代码结构详解

### **数据结构定义**

```go
// 双向链表节点
type Node struct {
    key   int
    value int
    prev  *Node  // 指向前一个节点
    next  *Node  // 指向后一个节点
}

// 双向链表
type DoubleLinkedList struct {
    head *Node  // 链表头部（最近使用）
    tail *Node  // 链表尾部（最久未使用）
    size int    // 当前大小
}

// LRU缓存
type LRUCache struct {
    capacity     int               // 缓存容量
    cache        map[int]*Node     // 哈希表：快速查找
    linkedList   *DoubleLinkedList // 双向链表：维护顺序
}
```

### **核心链表操作**

#### 1. 添加到头部
```go
  func (dll *DoubleLinkedList) addToHead(node *Node) {
      // 第1步：设置新节点的前驱为nil（头部节点没有前驱）
      node.prev = nil
      // 第2步：将新节点的后继指向当前头部
      node.next = dll.head

      // 第3步：检查当前链表是否为空
      if dll.head != nil {
          // 第4步：如果不为空，让原头部的前驱指向新节点
          dll.head.prev = node
      } else {
          // 第5步：如果为空，新节点既是头部也是尾部
          dll.tail = node
      }

      // 第6步：更新链表头部为新节点
      dll.head = node

      // 第7步：链表大小加1
      dll.size++
  }
```

#### 2. 移除指定节点
```go
func (dll *DoubleLinkedList) removeNode(node *Node) {
    // 处理前驱节点
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        // 移除的是头节点
        dll.head = node.next
    }

    // 处理后继节点
    if node.next != nil {
        node.next.prev = node.prev
    } else {
        // 移除的是尾节点
        dll.tail = node.prev
    }

    dll.size--
}
```

#### 3. 移动到头部（访问时使用）
```go
func (dll *DoubleLinkedList) moveToHead(node *Node) {
    dll.removeNode(node)   // 先从原位置移除
    dll.addToHead(node)    // 再添加到头部
}
```

#### 4. 移除尾部（缓存满时使用）
```go
func (dll *DoubleLinkedList) removeTail() *Node {
    if dll.tail == nil {
        return nil
    }

    tailNode := dll.tail
    dll.removeNode(tailNode)
    return tailNode
}
```

## 📊 时间复杂度分析

### **Get操作**：O(1)
- 哈希表查找：O(1)
- 移动到头部：O(1)
- 总计：O(1)

### **Put操作**：O(1)
- 哈希表查找/插入：O(1)
- 链表操作：O(1)
- 总计：O(1)

### **空间复杂度**：O(n)
- 哈希表存储：O(n)
- 双向链表存储：O(n)

## 🎮 实际应用场景

### **1. 浏览器缓存**
- 最近访问的网页保存在内存中
- 内存不足时，删除最久未访问的网页

### **2. 数据库查询缓存**
- 缓存最近查询的结果
- 提高重复查询的响应速度

### **3. 操作系统页面置换**
- 内存管理中的重要算法
- 物理内存不足时，选择页面换出到磁盘

### **4. Redis等缓存系统**
- 作为内存淘汰策略之一
- 常用于热点数据管理

## 🔄 与其他缓存策略对比

| 策略 | 优点 | 缺点 | 适用场景 |
|------|------|------|----------|
| **LRU** | 符合局部性原理，效率高 | 实现复杂度高 | 通用场景 |
| **FIFO** | 实现简单 | 可能删除热点数据 | 简单应用 |
| **LFU** | 考虑访问频率 | 实现更复杂 | 需要统计频率的场景 |
| **Random** | 实现简单 | 效果不稳定 | 特定场景 |

## 💡 核心思想总结

```
LRU = 哈希表 + 双向链表
哈希表：快速查找  O(1)
双向链表：维护顺序  O(1)

核心思想：
- 新数据放在头部（最近使用）
- 访问过的数据移到头部（更新时间）
- 容量不足时删除尾部数据（最久未使用）

时间复杂度：Get和Put都是O(1)
空间复杂度：O(n)

优势：
- 高效的访问速度
- 合理的淘汰策略
- 广泛的应用场景
```

## 🚀 常见问题与优化

### **Q: 为什么不用数组实现？**
A: 数组删除/插入需要O(n)时间，而双向链表只需O(1)

### **Q: 可以用单向链表吗？**
A: 理论上可以，但删除操作需要O(n)时间，性能下降

### **Q: 哈希表冲突怎么办？**
A: Go语言的map使用高效的哈希算法，冲突概率很低

### **Q: 大容量如何优化？**
A: 可以考虑分片、并发控制等技术

**掌握LRU缓存，你就掌握了数据结构与算法的经典应用！**🎉