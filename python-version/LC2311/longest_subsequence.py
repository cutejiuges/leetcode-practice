class Solution:
    @staticmethod
    def longestSubsequence(s: str, k: int) -> int:
        n, m = len(s), k.bit_length()
        if n < m:
            return n
        suf_val = int(s[n-m:], 2)
        ans = m - 1 if suf_val > k else m
        return ans + s[:n-m].count('0')
