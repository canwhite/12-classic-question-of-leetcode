# 12 Classic LeetCode Questions

Comprehend by Analogy

---

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

---

**View in Chinese**: [README_CN.md](README_CN.md)