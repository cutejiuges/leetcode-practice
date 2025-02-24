class Solution:
    def similarPairs(self, words: list[str]) -> int:
        ans = 0
        map = {}
        for ss in words:
            mask = 0
            for c in ss:
                mask |= 1 << (ord(c) - ord('a'))
            cnt = map.get(mask, 0)
            ans += cnt
            map[mask] = cnt+1
        return ans