class Solution:
    def percentageLetter(self, s: str, letter: str) -> int:
        cnt, n = 0, len(s)
        for ch in s:
            if ch == letter:
                cnt += 1
        return cnt * 100 // n
