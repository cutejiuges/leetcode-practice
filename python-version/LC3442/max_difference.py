from math import inf


class Solution:
    @staticmethod
    def maxDifference(s: str) -> int:
        cnt = [0] * 26
        for ch in s:
            cnt[ord(ch) - ord('a')] += 1
        mn, mx = inf, -inf
        for num in cnt:
            if num == 0:
                continue
            if num & 1 == 1:
                mx = max(mx, num)
            else:
                mn = min(mn, num)
        return mx - mn
