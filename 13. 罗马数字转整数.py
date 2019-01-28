class Solution:
    def romanToInt(self, s):
        """
        :type s: str
        :rtype: int
        """
        a = {'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000}
        length = len(s)
        ret = 0
        for i,v in enumerate(s):
            val = a[v] 
            if i + 1 < length:
                v2 = s[i+1]
                val2 = a[v2]
                if val2 > val:
                    val *= -1
            ret += val
        return ret
