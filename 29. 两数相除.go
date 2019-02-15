func divide(dividend int, divisor int) int {
    var dividendAbs, divisorAbs int = dividend, divisor
    if dividend < 0{
        dividendAbs = dividend * -1
    }
    if divisor < 0{
        divisorAbs = divisor * -1
    }
    remainder := dividendAbs
    ret := 0
    for remainder >= divisorAbs {
        remainder -= divisorAbs
        ret += 1
    }
    if dividend < 0{
        ret *= -1
    } 
    if divisor < 0 {
        ret *= -1
    }
    if ret == 1 << 31{
        ret = 1 << 31 - 1
    }
    return ret
}
