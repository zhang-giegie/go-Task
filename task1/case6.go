package task1
func RemoveDuplicates(nums []int) int {
     if len(nums) == 0 {
        return 0
    }

    j := 0 // j指针用于记录不重复元素的位置

    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[j] {
            j++
            nums[j] = nums[i]
        }
    }

    return j + 1 // 返回不重复元素的个数
}