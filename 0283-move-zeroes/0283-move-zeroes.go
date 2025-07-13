func moveZeroes(nums []int)  {
    zeroIndex := 0
    n := len(nums)

    for nonZeroIndex := 0; nonZeroIndex < n; nonZeroIndex++ {
        if nums[nonZeroIndex] != 0 {
            nums[zeroIndex] = nums[nonZeroIndex]
            zeroIndex++
        }
    }

    for i := zeroIndex; i < n; i++ {
        nums[i] = 0
    }
}