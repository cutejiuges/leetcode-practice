class Solution:
    def minimizedStringLength(self, s: str) -> int:
        mask = 0
        for ch in s:
            mask |= 1 << (ord(ch) - ord('a'))
        return mask.bit_count()
