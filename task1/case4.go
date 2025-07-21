package task1

func LongestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    // 以第一个字符串为基准
    prefix := strs[0]

    for i := 1; i < len(strs); i++ {
        // 比较当前字符串与基准前缀
        for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
            if prefix[j] != strs[i][j] {
                // 如果字符不匹配，截断前缀
                prefix = prefix[:j]
                break
            }
        }
        // 如果当前字符串比前缀短，更新前缀
        if len(strs[i]) < len(prefix) {
            prefix = prefix[:len(strs[i])]
        }
    }

    return prefix
}