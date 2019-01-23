class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        str_x = str(x)
        if str_x.startswith("-"):
            ret =  int("-"+str_x[1:][::-1])
        else:
            ret = int(str_x[::-1])
        if abs(ret) > (1 << 32) / 2:
            return 0 
        return ret
