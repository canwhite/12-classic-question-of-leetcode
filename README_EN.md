# 12 Classic LeetCode Questions

Comprehend by Analogy

<!-- Language Selection -->
<div align="right">
  <small>
    Language: 
    <strong>English</strong> | 
    <a href="#chinese" onclick="showLanguage('chinese'); return false;">中文</a>
  </small>
</div>

<!-- English Content -->
<div id="english-content">

### **The Ultimate 12 Classic Questions List**

1.  **1. Two Sum**

    - **Core Concept**: Hash Table
    - **Link**: https://leetcode.com/problems/two-sum/

2.  **206. Reverse Linked List**

    - **Core Concept**: Linked List Operations
    - **Link**: https://leetcode.com/problems/reverse-linked-list/

3.  **20. Valid Parentheses**

    - **Core Concept**: Stack
    - **Link**: https://leetcode.com/problems/valid-parentheses/

4.  **21. Merge Two Sorted Lists**

    - **Core Concept**: Linked List, Recursion/Iteration
    - **Link**: https://leetcode.com/problems/merge-two-sorted-lists/

5.  **70. Climbing Stairs**

    - **Core Concept**: Dynamic Programming Introduction
    - **Link**: https://leetcode.com/problems/climbing-stairs/

6.  **121. Best Time to Buy and Sell Stock**

    - **Core Concept**: Dynamic Programming/Greedy
    - **Link**: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

7.  **104. Maximum Depth of Binary Tree**

    - **Core Concept**: Binary Tree, DFS/BFS
    - **Link**: https://leetcode.com/problems/maximum-depth-of-binary-tree/

8.  **94. Binary Tree Inorder Traversal**

    - **Core Concept**: Binary Tree, Iteration/Recursion
    - **Link**: https://leetcode.com/problems/binary-tree-inorder-traversal/

9.  **3. Longest Substring Without Repeating Characters**

    - **Core Concept**: Sliding Window
    - **Link**: https://leetcode.com/problems/longest-substring-without-repeating-characters/

10. **146. LRU Cache**

    - **Core Concept**: Data Structure Design (Hash Table + Doubly Linked List)
    - **Link**: https://leetcode.com/problems/lru-cache/

11. **42. Trapping Rain Water**

    - **Core Concept**: Two Pointers, Dynamic Programming, Monotonic Stack
    - **Link**: https://leetcode.com/problems/trapping-rain-water/

12. **53. Maximum Subarray**
    - **Core Concept**: Dynamic Programming (Key Technique)
    - **Link**: https://leetcode.com/problems/maximum-subarray/

---

### **Why These 12 Questions? What Core Concepts Do They Reveal?**

These 12 questions together build a complete algorithmic ability framework, revealing four core aspects:

#### **Aspect 1: Data Structures are the Foundation of Algorithms (Questions: 1, 20, 21, 206, 146)**

- **Core Revelation**: Efficient algorithm implementation depends on selecting and skillfully operating appropriate data structures.
- **1 (Hash Table)**: Demonstrates how to use **space for time**, achieving fast lookups.
- **20 (Stack)**: Shows how to handle problems with **"recent relevance"**.
- **21 & 206 (Linked List)**: Shows that **pointer/reference operations** are everything in linked list algorithms, serving as micro-models for understanding recursion and iteration.
- **146 (Design)**: Demonstrates how to **combine data structures** (hash table + doubly linked list) to meet complex performance requirements, which is the core idea of system design.

#### **Aspect 2: The Starting Point and Prototype of Algorithmic Thinking (Questions: 70, 94, 104, 121)**

- **Core Revelation**: All complex algorithms evolve from the combination of basic ideas (recursion, divide and conquer, DP, greedy).
- **70 (DP)**: Reveals the most essential elements of dynamic programming - **state definition** and **state transition equations**.
- **94 & 104 (Tree)**: Reveals that **tree traversal** is the foundation of all tree problems, with recursive writing being an intuitive embodiment of **divide and conquer** thinking, while iterative writing profoundly reveals the relationship between **DFS/BFS** and **stack/queue**.
- **121 (DP/Greedy)**: Reveals the multi-dimensional expansion of dynamic programming states, while contrasting the simplicity of **greedy** algorithms for specific problems.

#### **Aspect 3: Powerful Tools for Efficient Problem Solving (Questions: 3, 53)**

- **Core Revelation**: Mastering one core technique can solve a large category of high-frequency interview questions.
- **3 (Sliding Window)**: Reveals how to use **sliding window** to efficiently handle **substring/subarray** problems, with the core being maintaining window validity.
- **53 (Dynamic Programming)**: Reveals an extremely important technique in dynamic programming - **"Kadane's Algorithm"**. Its state definition `dp[i]` is not the result directly asked by the problem, but rather **the maximum sum of subarrays ending with `nums[i]`**. This way of "defining subproblem states" is the key to solving many complex DP problems.

#### **Aspect 4: The Touchstone of Thinking Depth (Question: 42)**

- **Core Revelation**: Top-tier questions test the ability to integrate knowledge and solve problems from multiple angles, not just single knowledge points.
- **42 (Trapping Rain Water)**: This is a "methodology" question that reveals solutions from at least three perspectives:
  1.  **Dynamic Programming**: Precompute information, trade space for time.
  2.  **Two Pointers**: Extreme optimization of the DP method, directly calculating results during left and right pointer movement, with O(1) space complexity.
  3.  **Monotonic Stack**: Use a different data structure (stack) to track indices that might form "depressions", with completely different thinking.
- **The core it reveals is: True algorithmic ability is not about memorizing templates, but about flexibly applying and combining learned tools according to problem characteristics.**

### **Summary**

The value of this list lies in its **representativeness and efficiency**:

- **Representativeness**: Each question is a "knowledge anchor" - master one, and you can solve a category of problems by analogy.
- **Efficiency**: It covers over 90% of the core test points in interviews (arrays, linked lists, hash tables, stacks, queues, trees, dynamic programming, greedy, two pointers, sliding windows, design).

Through these 12 questions, you build not an isolated question bank, but an **interconnected, organic algorithmic knowledge system**.

</div>

<!-- Chinese Content -->
<div id="chinese-content" style="display: none;">

### **十二道终极经典题清单**

1.  **1. 两数之和** (Two Sum)

    - **核心考点**：哈希表
    - **链接**: https://leetcode-cn.com/problems/two-sum/

2.  **206. 反转链表** (Reverse Linked List)

    - **核心考点**：链表操作
    - **链接**: https://leetcode-cn.com/problems/reverse-linked-list/

3.  **20. 有效的括号** (Valid Parentheses)

    - **核心考点**：栈
    - **链接**: https://leetcode-cn.com/problems/valid-parentheses/

4.  **21. 合并两个有序链表** (Merge Two Sorted Lists)

    - **核心考点**：链表、递归/迭代
    - **链接**: https://leetcode-cn.com/problems/merge-two-sorted-lists/

5.  **70. 爬楼梯** (Climbing Stairs)

    - **核心考点**：动态规划入门
    - **链接**: https://leetcode-cn.com/problems/climbing-stairs/

6.  **121. 买卖股票的最佳时机** (Best Time to Buy and Sell Stock)

    - **核心考点**：动态规划/贪心
    - **链接**: https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/

7.  **104. 二叉树的最大深度** (Maximum Depth of Binary Tree)

    - **核心考点**：二叉树、DFS/BFS
    - **链接**: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/

8.  **94. 二叉树的中序遍历** (Binary Tree Inorder Traversal)

    - **核心考点**：二叉树、迭代/递归
    - **链接**: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

9.  **3. 无重复字符的最长子串** (Longest Substring Without Repeating Characters)

    - **核心考点**：滑动窗口
    - **链接**: https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

10. **146. LRU 缓存** (LRU Cache)

    - **核心考点**：数据结构设计（哈希表+双向链表）
    - **链接**: https://leetcode-cn.com/problems/lru-cache/

11. **42. 接雨水** (Trapping Rain Water)

    - **核心考点**：双指针、动态规划、单调栈
    - **链接**: https://leetcode-cn.com/problems/trapping-rain-water/

12. **53. 最大子数组和** (Maximum Subarray)
    - **核心考点**：动态规划（关键技巧）
    - **链接**: https://leetcode-cn.com/problems/maximum-subarray/

---

### **为什么是这十二道？它们揭示了什么核心？**

这十二道题共同构建了一个完整的算法能力框架，揭示了四大核心层面：

#### **层面一：数据结构是算法的基石 (题号: 1, 20, 21, 206, 146)**

- **核心揭示**：算法的高效实现依赖于选择并熟练操作合适的数据结构。
- **1 (哈希表)**：展示了如何用**空间换时间**，实现快速查找。
- **20 (栈)**：展示了如何处理具有**"最近相关性"**的问题。
- **21 & 206 (链表)**：展示了**指针/引用操作**是链表算法的全部，是理解递归和迭代的微观模型。
- **146 (设计)**：展示了如何**组合数据结构**（哈希表+双向链表）来满足复杂的性能需求，这是系统设计的核心思想。

#### **层面二：算法思想的起点与雏形 (题号: 70, 94, 104, 121)**

- **核心揭示**：所有复杂算法都是由基础思想（递归、分治、DP、贪心）组合演变而来。
- **70 (DP)**：揭示了动态规划最本质的要素——**状态定义**和**状态转移方程**。
- **94 & 104 (树)**：揭示了**树的遍历**是所有树问题的基础，其递归写法是**分治思想**的直观体现，迭代写法则深刻揭示了**DFS/BFS**与**栈/队列**的关系。
- **121 (DP/贪心)**：揭示了动态规划状态的多维度扩展，同时对比了**贪心**算法在特定问题上的简洁性。

#### **层面三：高效解决问题的利器 (题号: 3, 53)**

- **核心揭示**：掌握一种核心技巧，就能解决一大类高频面试题。
- **3 (滑动窗口)**：揭示了如何用**滑动窗口**高效处理**子串/子数组**问题，核心在于维护窗口的合法性。
- **53 (动态规划)**：揭示了动态规划中一个极其重要的技巧——**"Kadane 算法"**。它的状态定义 `dp[i]` 不是问题直接要问的结果，而是**以 `nums[i]` 为结尾的子数组的最大和**。这种"定义子问题状态"的方式是解决许多复杂 DP 问题的钥匙。

#### **层面四：思维深度的试金石 (题号: 42)**

- **核心揭示**：顶尖的题目考察的是融会贯通、多角度解决问题的能力，而非单一知识点。
- **42 (接雨水)**：这是一道"方法论"题目，它从至少三个角度揭示了解决方案：
  1.  **动态规划**：预计算信息，用空间换时间。
  2.  **双指针**：对 DP 方法的极致优化，在左右指针的移动中直接计算结果，空间复杂度 O(1)。
  3.  **单调栈**：换一种数据结构（栈）来跟踪可能形成"洼地"的索引，思维完全不同。
- **它揭示的核心是：真正的算法能力不是背诵模板，而是根据问题特征，灵活运用和组合所学工具的能力。**

### **总结**

这份清单的价值在于它的**代表性和效率**：

- **代表性**：每一道题都是一个"知识锚点"，掌握一道，便能触类旁通，解决一类问题。
- **效率**：它覆盖了面试中超过 90%的核心考点了（数组、链表、哈希表、栈、队列、树、动态规划、贪心、双指针、滑动窗口、设计）。

通过这十二道题，你构建起的不是一个孤立的题库，而是一个**相互关联的、有机的算法知识体系**。

</div>

<!-- JavaScript for Language Switching -->
<script>
function showLanguage(lang) {
    if (lang === 'english') {
        document.getElementById('english-content').style.display = 'block';
        document.getElementById('chinese-content').style.display = 'none';
        // Update language selector
        document.querySelector('a[href="#english"]').parentNode.innerHTML = '<strong>English</strong>';
        document.querySelector('a[href="#chinese"]').parentNode.innerHTML = ' | <a href="#chinese" onclick="showLanguage(\'chinese\'); return false;">中文</a>';
    } else if (lang === 'chinese') {
        document.getElementById('english-content').style.display = 'none';
        document.getElementById('chinese-content').style.display = 'block';
        // Update language selector
        document.querySelector('a[href="#english"]').parentNode.innerHTML = '<a href="#english" onclick="showLanguage(\'english\'); return false;">English</a>';
        document.querySelector('a[href="#chinese"]').parentNode.innerHTML = ' | <strong>中文</strong>';
    }
}

// Set English as default
window.onload = function() {
    showLanguage('english');
};
</script>