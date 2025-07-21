package task1

import (
    "sort"
)

func Merge(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }

    // 按照区间的起始值进行排序
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })

    merged := [][]int{}
    for _, interval := range intervals {
        // 如果结果数组为空或者当前区间的起始值大于结果数组中最后一个区间的结束值
        if len(merged) == 0 || interval[0] > merged[len(merged)-1][1] {
            merged = append(merged, interval)
        } else {
            // 否则，合并区间
            merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
        }
    }

    return merged
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
