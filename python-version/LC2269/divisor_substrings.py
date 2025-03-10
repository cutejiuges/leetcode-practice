class Solution:
    def divisorSubstrings(self, num: int, k: int) -> int:
        ss = str(num)
        n = len(ss)
        ans = 0
        for i in range(n-k+1):
            t = int(ss[i:i+k])
            if t != 0 and num % t == 0:
                ans += 1
        return ans
