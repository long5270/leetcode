func strStr(haystack string, needle string) int {
    runeNeedle := []rune(needle) 
    runeHaystack := []rune(haystack)
    lenght1 := len(runeNeedle)
    lenght2 := len(runeHaystack)
    if lenght1 == 0{
        return 0
    }
    var ret int = -1
    var flag int
    for i:=0; i < lenght2; i++ {
        if runeHaystack[i] == runeNeedle[flag]{
            flag += 1
            if flag == lenght1{
                ret = i - flag + 1
                break
            }
        }else{
            if flag > 0 {
                i = i - flag
                flag = 0
            }
        }
    }
    return ret
}
