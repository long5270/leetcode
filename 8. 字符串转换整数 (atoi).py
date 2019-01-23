class Solution:
    def myAtoi(self, s):
        """
        :type s: str
        :rtype: int
        """
        ret = ""
        s = s.lstrip()
        if s.startswith("-") or s.startswith("+"):
            ret = s[0]
            s = s[1:]
        for i in s:
            if i.isdigit():
                ret += i
            else:
                break
        if ret == "-" or ret == "+" or not ret:
            return 0
        int_ret = int(ret)
        if abs(int_ret) >= 1 << 31:
            if int_ret > 0:
                return (1 << 31) - 1
            else:
                return (1 << 31) * -1
        return int_ret
