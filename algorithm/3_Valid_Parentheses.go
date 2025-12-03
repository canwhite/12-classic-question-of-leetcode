package algorithm

// 有效括号，栈的应用，用栈做匹配
// 这里需要注意的是栈是先进后出的，如果外层是中括号，里边是小括号，我们需要比对小括号
// 单调栈虽然主要用在接雨水等问题，但括号匹配也体现了栈的核心思想
func IsValid(s string) bool {
    // 使用栈来匹配括号（这里体现了栈的LIFO特性）
    stack := make([]rune, 0)

    // 定义括号映射关系，右括号对应左括号
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s {
        // 如果是左括号，入栈（等待匹配）
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else if char == ')' || char == '}' || char == ']' {
            // 如果是右括号，检查栈是否为空或栈顶不匹配
            if len(stack) == 0 {
                return false
            }
            // 拿出栈顶（最近的左括号）
            top := stack[len(stack)-1]

            // 匹配，不匹配返回false
            if top != pairs[char] {
                return false
            }
            // 匹配成功，出栈
            stack = stack[:len(stack)-1]
        }
    }

    // 最终栈为空说明所有括号都匹配
    return len(stack) == 0
}

// 使用哈希表的优化版本（更简洁）
func IsValidOptimized(s string) bool {
    stack := make([]rune, 0)
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s {
        // 左括号入栈
        if pairs[char] == 0 { // 不在pairs中说明是左括号
            stack = append(stack, char)
        } else { // 右括号，需要匹配
            if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}

// 不使用额外映射关系的版本（更高效）
func IsValidEfficient(s string) bool {
    stack := make([]byte, 0)

    for i := 0; i < len(s); i++ {
        char := s[i]
        switch char {
        case '(', '{', '[':
            stack = append(stack, char)
        case ')':
            if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            stack = stack[:len(stack)-1]
        case '}':
            if len(stack) == 0 || stack[len(stack)-1] != '{' {
                return false
            }
            stack = stack[:len(stack)-1]
        case ']':
            if len(stack) == 0 || stack[len(stack)-1] != '[' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
    }

    return len(stack) == 0
}