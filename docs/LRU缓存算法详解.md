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

### **双向链表 vs 单向链表对比**

| 操作 | 双向链表 | 单向链表 | 说明 |
|------|----------|----------|------|
| **删除节点** | O(1) | O(n) | 双向链表通过prev指针直接删除 |
| **移动节点** | O(1) | O(n) | 双向链表断开重连，单向需要遍历 |
| **插入头部** | O(1) | O(1) | 两者都可以 |
| **获取前驱** | O(1) | 不可能 | 单向链表无法向前遍历 |
| **内存占用** | 稍多 | 较少 | 双向需要prev指针 |

### **核心差异详解**

#### **双向链表删除操作示例**
```go
// 删除任意节点，这种事简单说法，当然有些时候还要考虑为头尾的情况
node.prev.next = node.next  // 通过prev直接修改前驱的next
node.next.prev = node.prev  // 通过next直接修改后继的prev

```

#### **单向链表删除操作示例**
```go
// 删除节点需要先找到前驱
func findPrev(head, node *Node) *Node {
    current := head
    for current != nil && current.next != node {
        current = current.next  // 需要遍历 O(n)
    }
    return current
}

// 删除操作
prev := findPrev(dll.head, node)  // O(n)
prev.next = node.next              // O(1)

```

#### 双向链表 vs 单向链表插入对比

**双向链表插入到头部**：
```go
func (dll *DoubleLinkedList) addToHead(node *Node) {
    // === 节点级沟通：节点之间的双向通知 ===

    // 第1步：新节点声明自己前面没人（作为新队首）
    node.prev = nil

    // 第2步：新节点记录现在队首是谁（准备插队）
    node.next = dll.head

    // === 队列级沟通：更新队列管理信息 ===

    //第三步： 也是队列节点先沟通
    if dll.head != nil {
        // 第4步：通知现有队首"你现在有上级了"
        dll.head.prev = node
    } else {
        // 第5步：空队伍时，新节点既是队首也是队尾
        dll.tail = node
    }


    // 第6步：官方宣布新队首
    dll.head = node

    // 第7步：更新队伍人数
    dll.size++
}
```

**单向链表插入到头部**：
```go
func addToHead(node *Node, head **Node) {
    // 就这两步，很简单！
    node.next = *head        // 新节点指向原头
    *head = node            // 更新头指针

    // 就这么简单！没有反向指针的烦恼
}
```

**🎯 对比总结**：
- **双向链表**：需要处理4个指针连接，复杂但功能强大
- **单向链表**：只需要处理2个指针，简单高效
- **删除操作才是关键区别**：双向O(1) vs 单向O(n)

### **双向链表的核心价值**

**1. 两头连接的能力**
```go
// 每个节点都能向前和向后操作
node ←→ node ←→ node
  ↑ prev    ↑ next
```

**2. 任意位置O(1)操作**
- 不需要从头部遍历
- 直接通过节点自身的前后指针操作
- 真正实现LRU的O(1)性能

**3. 灵活的移动能力**
```go
// LRU中常见操作：
this.linkedList.removeNode(node)    // 从任意位置删除 O(1)
this.linkedList.addToHead(node)      // 添加到头部 O(1)
this.linkedList.removeTail()         // 删除尾部 O(1)
```

### **为什么单向链表无法高效实现LRU？**

**时间复杂度对比**：
| 操作 | 双向链表LRU | 单向链表LRU |
|------|-------------|-------------|
| Get(访问数据) | O(1) | O(n) |
| Put(添加数据) | O(1) | O(n) |
| 淘汰数据 | O(1) | O(n) |

**根本原因**：
- **双向链表**：每个节点都认识自己的前后邻居
- **单向链表**：只认识后一个邻居，要找前一个需要全村通知

**结论**：LRU缓存的高效性完全依赖于双向链表的两头连接能力！没有双向链表，就没有真正意义上的O(1)LRU缓存。

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
        // 步骤2：存在，将节点移到头部（标记为最近使用），头部就是最近使用
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
    size int    // 当前大小，双向链表的长度
}

// LRU缓存
type LRUCache struct {
    capacity     int               // 缓存容量
    cache        map[int]*Node     // 哈希表：快速查找
    linkedList   *DoubleLinkedList // 双向链表：维护顺序
}
```

### **核心链表操作**

#### 1. 添加到头部 - 两层沟通模式

**核心思想**：先让节点沟通，然后和队列沟通，这就是双向链表的精髓

```go
func (dll *DoubleLinkedList) addToHead(node *Node) {
    // === 节点级沟通：节点之间的双向通知 ===

    // 第1步：新节点声明自己前面没人（作为新队首）
    node.prev = nil

    // 第2步：新节点记录现在队首是谁（准备插队）
    node.next = dll.head

    // === 队列级沟通：更新队列管理信息 ===

    //第三步： 也是队列节点先沟通
    if dll.head != nil {
        // 第4步：通知现有队首"你现在有上级了"
        dll.head.prev = node
    } else {
        // 第5步：空队伍时，新节点既是队首也是队尾
        dll.tail = node
    }


    // 第6步：官方宣布新队首
    dll.head = node

    // 第7步：更新队伍人数
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