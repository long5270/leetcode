class Solution:
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if not strs:
            return ""
        if not strs[1:]:
            return strs[0]
        str1 = strs[0]
        flag = False
        ret = ''
        for i,v in enumerate(str1):
            for v2 in strs[1:]:
                if not v2[i:] or v2[i] != v:
                    flag = True
                    break
            if flag:
                break
            ret += v
        return ret
