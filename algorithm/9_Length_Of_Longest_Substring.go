package algorithm

//也可以理解为求解无重复字符串
func LengthOfLongestSubstring(s string) int {
    // 滑动窗口算法
    // 用 map 记录字符最后一次出现的位置，注意记录目的
	// 典型的用空间解决时间问题，三方记录
    charIndex := make(map[byte]int)

    // 窗口的左右边界
    left := 0
    maxLength := 0
 
    // 右边界自动移动
    for right := 0; right < len(s); right++ {
		
        currentChar := s[right]

        // 通过记忆包更新做节点
        if lastPos, exists := charIndex[currentChar]; exists && lastPos >= left {
            // 如果出现过，并且位置在当前窗口内
            // 需要将窗口左边界移动到重复字符的下一个位置
            left = lastPos + 1
        }

        // 记录当前字符的位置
        charIndex[currentChar] = right

        // 更新最大长度
        currentLength := right - left + 1
        if currentLength > maxLength {
            maxLength = currentLength
        }
    }

    return maxLength
}