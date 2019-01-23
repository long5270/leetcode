class Solution:
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        str_x = str(x)
        lenght = len(str_x)
        half = int(lenght / 2)
        if str_x[:half] == str_x[::-1][:half]:
            return True
        return False
