class Solution:
    def intToRoman(self, num):
        """
        :type num: int
        :rtype: str
        """
        ret = ""
        a = {1:"I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M"}
        b = [1000, 100, 10, 1]
        for i in b:
            tmpa = int(num / i)
            tmpb = num % i
            num = tmpb
            if i < 1000:
                if tmpa == 4:
                    ret += a[i] + a[i * 5] 
                    continue
                elif tmpa == 9:
                    ret = ret + a[i] + a[i * 10] 
                    continue
                elif tmpa >= 5:
                    ret += a[i*5]
                    tmpa -= 5
            ret += a[i] * tmpa
            if tmpb == 0:
                break
            num = tmpb
        return ret
