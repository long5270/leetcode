func twoSum(nums []int, target int) []int {
    for i,iv := range nums{
        for j,jv := range nums{
            if j!=i && iv + jv == target{
                return []int{i,j}
            }
        }
    }
    return []int{}
}
