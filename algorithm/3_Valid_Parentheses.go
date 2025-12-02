package algorithm

//有效括号，栈的应用，用栈做匹配，
//这里需要注意的是栈是先进后出的，如果外层是中括号，里边是小括号，我们回西安比对小括号
func IsValid(s string) bool {
    // 使用栈来匹配括号
    stack := make([]rune, 0)

    // 定义括号映射关系
    pairs := map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    for _, char := range s {
        // 如果是左括号，入栈
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        } else if char == ')' || char == '}' || char == ']' {
            // 如果是右括号，检查栈是否为空或栈顶不匹配
            if len(stack) == 0 {
                return false
            }
			//拿出栈顶
            top := stack[len(stack)-1]
			//匹配，不能成返回false
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