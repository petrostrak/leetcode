func twoSum(nums []int, target int) []int {
    seen := make(map[int]int)
    var arr []int
    for i := range nums {
        diff := target - nums[i]
        if _, ok := seen[diff]; !ok {
            seen[nums[i]] = i
        } else {
            arr = append(arr, seen[diff])
            arr = append(arr, i)
        }
    }
    return arr
}