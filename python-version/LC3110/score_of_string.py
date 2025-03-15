class Solution:
    def scoreOfString(self, s: str) -> int:
        ans = 0
        n = len(s)
        for i, ch in enumerate(s):
            if i < n-1:
                ans += abs(ord(ch) - ord(s[i+1]))
        return ans
