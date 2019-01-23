class Solution:
    def convert(self, s, numRows):
        """
        :type s: str
        :type numRows: int
        :rtype: str
        """
        if numRows == 1:
            return s
        rows = []
        for i in range(numRows):
            rows.append([])
        step = numRows * 2 - 2
        for i in range(len(s)):
            tmp = i % step
            if tmp < numRows:
                rows[tmp].append(s[i])
            else:
                rows[1:-1][numRows-3-(tmp - numRows)].append(s[i])
        return "".join(["".join(row) for row in rows])
