# DFS深度优先搜索 - 小白从零开始完全指南

## 🎯 什么是DFS？超级简单理解！

### 一句话理解：
**DFS = 一条路走到黑，然后换条路继续走**    
当然需要了解的一点是，第一层是遍历的，然后往下钻，就是我们在做的    


### 生活比喻：你逛大型商场找钥匙

```
商場地圖：
入口 → 大廳A → 走廊B → 商店C → 走廊D → 商店E
        ↓        ↓        ↓        ↓
     找鑰匙   走廊F   商店G → 走廊H → 商店Iwo
              ↓        ↓        ↓        ↓
             找鑰匙   走廊J → 出口
```

**你的逛街策略**：
1. 从入口开始，选择**一条主路**走到底
2. 路过每一家商店都**进去看看**有没有钥匙
3. 如果走到**死胡同**（没有钥匙），就**退回到上一个岔路口**
4. 选择**另一条路**继续走
5. 重复直到找到钥匙或者逛完所有地方

## 🧠 DFS的核心思想

### 三大原则：

1. **一条路走到底**：深入探索，不轻易回头
2. **遇到死胡同才回头**：确认这条路走不通才回溯
3. **记住去过的地方**：避免走回头路重复探索

### 伪代码理解：
```
DFS(当前位置):
    if 当前位置是目标:
        找到了！返回成功
    if 当前位置已经去过:
        走过了，跳过
    标记当前位置为"已去过"

    for 每个可以去的方向:
        if 方向可以走:
            下一个位置 = 方向
            result = DFS(下一个位置)
            if result == "找到了":
                返回成功
        // 如果这个方向的所有路都走完了还没找到，自动回到上一层
    返回失败
```

## 🚀 DFS的两种实现方式

### 方式1：递归实现（最常用，最容易理解）
# 这里实际上还用了map提高效率，另类的额用空间换时间，map的value是我们的目的，防止重复

#### 核心代码：
```go
func DFS(start *Node, target interface{}, visited map[*Node]bool) *Node {
    // 1. 基础判断：当前位置是否有效
    if start == nil {
        return nil
    }

    // 2. 判断是否找到目标
    if start.Value == target {
        return start
    }

    // 3. 标记当前位置为"已访问"
    visited[start] = true

    // 4. 对每个方向进行探索
    for _, neighbor := range start.Neighbors {
        if !visited[neighbor] {
            // 递归进入下一个位置，继续深入！
            result := DFS(neighbor, target, visited)
            if result != nil {
                // 找到了！把结果一层一层传回去
                return result
            }
        }
    }

    // 5. 所有方向都试过了，没找到，返回失败
    return nil
}
```

#### 完整例子演示：

假设有这样一个迷宫图：
```
入口 → A → B → C → 目标
       │
       D → E
```

**执行过程（递归调用栈）**：
```
第1层：DFS(入口)
  标记：入口已访问
  探索：DFS(A)

第2层：DFS(A)
  标记：A已访问
  探索：DFS(B)

第3层：DFS(B)
  标记：B已访问
  探索：DFS(C)

第4层：DFS(C)
  标记：C已访问
  发现：找到了目标！
  返回：目标

第3层：收到"找到了！"，返回：C → 目标

第2层：收到"C → 目标"，返回：B → C → 目标

第1层：收到"B → C → 目标"，返回：A → B → C → 目标
```

### 递归的可视化：
```
函数调用栈（像一叠纸）：
┌─────────────────────────┐
│ 第4层：DFS(C)找到目标！│ ← 最上面
├─────────────────────────┤
│ 第3层：DFS(C)         │
├─────────────────────────┤
│ 第2层：DFS(B)         │
├─────────────────────────┤
│ 第1层：DFS(入口)       │ ← 最下面
└─────────────────────────┘

每层的"返回"就是"退回到上一层"的意思！
```

### 递归的优缺点：

#### 优点：
✅ **代码简洁**：符合人的思维直觉
✅ **逻辑清晰**：一条路走到底的思路很明显
✅ **自动回溯**：函数返回时自动"回到上一层"

#### 缺点：
❌ **可能栈溢出**：路径太长时，递归层数太多
❌ **内存消耗**：每次递归调用都需要额外的内存
❌ **性能风险**：不适合处理很深的图

---

### 方式2：非递归实现（用栈模拟递归）

#### 核心思想：
用**栈**来记录要走的路，手动模拟递归过程：
- **入栈**：发现一个新地方，加入待探索列表
- **出栈**：探索完一个地方，从栈顶取下一个地方
- **回溯**：自动退回到上一个地方

#### 核心代码：
```go
func DFSIterative(start *Node, target interface{}) *Node {
    if start == nil {
        return nil
    }

    // 1. 准备栈和访问记录
    stack := []*Node{start}  // 栈：要探索的地方（后进先出）
    visited := make(map[*Node]bool)
    visited[start] = true

    // 2. 只要还有地方要探索
    for len(stack) > 0 {
        // 取出栈顶（最后一个进来的，下一个要探索的）
        current := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // 检查是否找到目标
        if current.Value == target {
            return current
        }

        // 把当前地方的所有相邻地方都加入栈
        for _, neighbor := range current.Neighbors {
            if !visited[neighbor] {
                stack = append(stack, neighbor)
                visited[neighbor] = true
            }
        }
    }

    return nil
}
```

#### 手动过程演示：
```
初始：stack=[入口], visited=[入口]

第1步：
current = 入口
stack = []  (出栈)
把入口的邻居加入栈：[A, D]
visited = [入口, A, D]

第2步：
current = D  (栈顶)
stack = [A]   (出栈)
把D的邻居加入栈：[A, E]
visited = [入口, A, D, E]

第3步：
current = E
stack = [A]   (出栈)
E没有邻居，不加入
visited = [入口, A, D, E]

第4步：
current = A
stack = []   (出栈)
把A的邻居加入栈：[B]
visited = [入口, A, D, E, B]

第5步：
current = B
stack = []   (出栈)
把B的邻居加入栈：[C]
visited = [入口, A, D, E, B, C]

第6步：
current = C
stack = []   (出栈)
发现了目标！返回C
```

### 非递归的优缺点：

#### 优点：
✅ **不会栈溢出**：栈的大小受系统内存限制
✅ **内存可控**：可以精确控制内存使用
✅ **性能稳定**：适合处理很深的图

#### 缺点：
❌ **代码复杂**：需要手动管理栈
❌ **理解困难**：回溯过程不那么直观
❌ **调试复杂**：需要跟踪栈的状态变化

---

## 🎯 DFS的实际应用场景

### 1. 迷宫游戏求解
```go
// 找迷宫出口
func SolveMaze(start *MazeNode, end *MazeNode) []*MazeNode {
    visited := make(map[string]bool)
    return dfsMaze(start, end, visited, []*MazeNode{start})
}

func dfsMaze(current, end *MazeNode, visited map[string]bool, path []*MazeNode) []*MazeNode {
    if current == nil || current == end {
        if current == end {
            return path
        }
        return nil
    }

    visited[current.ID] = true
    path = append(path, current)

    // 尝试四个方向
    directions := []Direction{北, 东, 南, 西}
    for _, dir := range directions {
        next := current.GetNeighbor(dir)
        if next != nil && !visited[next.ID] {
            result := dfsMaze(next, end, visited, path)
            if result != nil {
                return result
            }
        }
    }

    // 回溯
    return nil
}
```

### 2. 文件系统遍历
```go
// 找特定文件
func FindFile(root *FileNode, filename string) *FileNode {
    return dfsFile(root, filename, make(map[string]bool))
}

func dfsFile(current *FileNode, filename string, visited map[string]bool) *FileNode {
    if current == nil {
        return nil
    }

    if current.Name == filename {
        return current
    }

    visited[current.Path] = true

    // 深度优先：先处理文件，再处理子目录
    if current.IsFile {
        for _, file := range current.Files {
            if !visited[file.Path] {
                result := dfsFile(file, filename, visited)
                if result != nil {
                    return result
                }
            }
        }
    }

    // 再处理子目录
    for _, dir := range current.Directories {
        if !visited[dir.Path] {
            result := dfsFile(dir, filename, visited)
            if result != nil {
                return result
            }
        }
    }

    return nil
}
```

### 3. 路径查找和环检测
```go
// 检测图中是否有环
func HasCycle(graph []*Node) bool {
    visited := make(map[*Node]bool)

    for _, node := range graph {
        if !visited[node] {
            if dfsCycle(node, visited, nil) {
                return true
            }
        }
    }
    }
    return false
}

func dfsCycle(node *Node, visited map[*Node]bool, parent *Node) bool {
    if visited[node] {
        return true  // 再次访问到，说明有环
    }

    visited[node] = true

    for _, neighbor := range node.Neighbors {
        if neighbor != parent {  // 避免误判
            if dfsCycle(neighbor, visited, node) {
                return true
            }
        }
    }

    return false
}
```

### 4. 拓扑排序
```go
// 课程依赖排序
func TopologicalSort(courses []*Course) []*Course {
    result := []*Course{}
    visited := make(map[int]bool)

    for _, course := range courses {
        if !visited[course.ID] {
            if dfsTopological(course, visited, &result) {
                return result
            }
        }
    }
    return nil
}

func dfsTopological(course *Course, visited map[int]bool, result *[]*Course) bool {
    if visited[course.ID] {
        return true
    }

    visited[course.ID] = true

    // 先处理所有依赖
    for _, prereq := range course.Prerequisites {
        if !visited[prereq.ID] {
            if !dfsTopological(prereq, visited, result) {
                return false  // 有环，无法排序
            }
        }
    }

    // 处理完依赖后，把自己加入结果
    *result = append(*result, course)
    return true
}
```

## 🎨 DFS的常见陷阱和解决方案

### 陷阱1：忘记标记已访问
```go
// ❌ 错误：会导致无限循环
func DFSCycle(start *Node) {
    if start == nil {
        return
    }
    fmt.Println("访问:", start.Value)

    for _, neighbor := range start.Neighbors {
        DFSCycle(neighbor)  // 可能又回到起点！
    }
}

// ✅ 正确：使用visited数组
func DFSCorrect(start *Node, visited map[*Node]bool) {
    if start == nil || visited[start] {
        return
    }

    fmt.Println("访问:", start.Value)
    visited[start] = true

    for _, neighbor := range start.Neighbors {
        if !visited[neighbor] {
            DFSCorrect(neighbor, visited)
        }
    }
}
```

### 陷阱2：递归太深导致栈溢出
```go
// ❌ 有风险的版本
func DFSDeep(start *Node) {
    // 如果链表很长，可能导致栈溢出
    if start == nil {
        return
    }
    return DFSDeep(start.Next)
}

// ✅ 优化：限制递归深度或使用迭代
func DFSLimited(start *Node, maxDepth int) {
    return dfsLimited(start, 0, maxDepth)
}

func dfsLimited(start *Node, currentDepth, maxDepth int) {
    if start == nil || currentDepth >= maxDepth {
        return
    }
    return dfsLimited(start.Next, currentDepth+1, maxDepth)
}
```

### 陷阱3：不正确的回溯处理
```go
// ❌ 错误：忘记清理状态
func DFSWrong(start, target interface{}, visited map[*Node]bool) *Node {
    visited[start] = true

    for _, neighbor := range start.Neighbors {
        if !visited[neighbor] {
            result := DFSWrong(neighbor, target, visited)
            if result != nil {
                return result
            }
            // ❌ 这里应该回溯visited状态！
        }
    }
    return nil
}

// ✅ 正确：或者使用不同的visited策略
func DFSCorrect(start *Node, target interface{}, visited map[*Node]bool, path []*Node) *Node {
    visited[start] = true
    path = append(path, start)

    for _, neighbor := range start.Neighbors {
        if !visited[neighbor] {
            result := DFSCorrect(neighbor, target, visited, path)
            if result != nil {
                return result
            }
            // ✅ path在每次递归调用时都是独立的副本
        }
    }

    return nil
}
```

## 💡 小白学习建议

### 第1步：理解"一条路走到底"
- 想象你在迷宫中，选择一条路坚持走
- 遇到死胡同才选择另一条路
- 这是DFS的核心思想

### 第2步：掌握递归的"自动回溯"
- 函数返回就是"退回到上一层"
- 不需要手动回溯，递归自动处理
- 理解函数调用栈的概念

### 第3步：学会使用visited数组
- 防止重复访问同一个地方
- 避免无限循环
- 这是DFS必须的安全措施

### 第4步：理解栈和递归的关系
- 递归隐含使用系统栈
- 非递归显式使用数据栈
- 两者本质相同，只是实现方式不同

### 第5步：练习实际应用
- 从简单的迷宫开始
- 再到文件系统遍历
- 最后到复杂的图问题

## 🎯 记忆口诀

### DFS三步曲：
```
1. 一条路走到底（深入探索）
2. 遇到死胡同才回（递归返回）
3. 记住去过的地方（避免重复）
```

### 选择DFS的时机：
- **需要任意一条路径**：迷宫求解、拓扑排序
- **内存敏感环境**：嵌入式系统
- **学习理解算法**：递归思想清晰

掌握了DFS，你就掌握了计算机科学中最重要的算法思想之一！当然如果你想，可以去了解下[BFS](./BFS小白从零开始详解.md)