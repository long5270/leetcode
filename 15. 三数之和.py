class Solution:
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        retset = set()
        nummap = {}
        for i, v in enumerate(nums):
            if v in nummap:
                nummap[v].append(i)
            else:
                nummap[v] = [i]
        length = len(nums)
        for i in range(length):
            for j in range(length):
                if i == j:
                    continue
                vi = nums[i]
                vj = nums[j]
                tmp = 0 - vi - vj
                tmplist = nummap.get(tmp)
                if not tmplist:
                    continue
                for q in tmplist:
                    if i == q or j == q:
                        continue
                    retset.add(tuple(sorted([vi, vj, tmp])))
                    break
        ret = []
        for r in retset:
            ret.append(list(r))
        return ret
