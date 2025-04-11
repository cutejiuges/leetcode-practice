class Solution:
    def countSymmetricIntegers(self, low: int, high: int) -> int:
        ans = 0
        for i in range(low, high+1):
            num = str(i)
            n = len(num)
            if n & 1 == 1:
                continue
            judge = 0
            for j in range(n//2):
                judge += ord(num[j])
            for j in range(n//2, n):
                judge -= ord(num[j])
            ans += 1 if judge == 0 else 0
        return ans
