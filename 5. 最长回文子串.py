class Solution:
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        len_s = len(s)
        if len_s < 2:
            return s
        r = s[::-1]
        ret = s[0]
        len_ret = 1
        for i in range(len_s):
            lenght = 1
            for j in range(len_s - i):
                while i + lenght <= len_s:
                    tmp = s[i: i + lenght]
                    if tmp == r[j: j + lenght]:
                        half = int(lenght / 2)
                        if tmp[:half] == tmp[::-1][:half]:
                            if lenght > len_ret:
                                len_ret = lenght
                                ret = tmp
                        lenght += 1
                    else:
                        break
        return ret
