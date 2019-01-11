func lengthOfLongestSubstring(s string) int {
    var max int
    var history map[string]int = make(map[string]int, 10)
    var current, beforeIndex int
    for i,v := range s{
        strV := string(v)
        tmp, ok := history[strV]
        if beforeIndex > tmp{
            ok = false
        }
        if ok{
            history[strV] = i
            if current > max {
                max = current
            }
            beforeIndex = tmp
            current = i - tmp
        }else{
            current += 1
            history[strV] = i
        }
    }
    if current > max {
        max = current
    }
    return max
    
}
