class Solution:
    @staticmethod
    def canBeTypedWords(text: str, broken_letters: str) -> int:
        mask = 0
        for s in broken_letters:
            mask |= 1 << (ord(s) - ord('a'))
        ans, ok = 0, 1
        for ch in text:
            if ch == ' ':
                ans += ok
                ok = 1
                continue
            if mask & (1 << ord(ch) - ord('a')) != 0:
                ok = 0
        ans += ok
        return ans
