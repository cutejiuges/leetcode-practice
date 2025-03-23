class Solution:
    def canBeValid(self, s: str, locked: str) -> bool:
        n = len(s)
        if n & 1 == 1:
            return False

        mn, mx = 0, 0
        for i, ch in enumerate(s):
            if locked[i] == '1':
                cur = 1 if ch == '(' else -1
                mx += cur
                if mx < 0:
                    return False
                mn += cur
            else:
                mx += 1
                mn -= 1
            mn = max(mn, 0)
        return mn == 0
