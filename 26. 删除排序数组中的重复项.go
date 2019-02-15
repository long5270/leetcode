func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    current := 0
    for _,v := range nums{
        if v > nums[current]{
            current += 1
            nums[current] = v
        }
    }
    return current + 1
}
