class Solution:
    def addSpaces(self, s: str, spaces: list[int]) -> str:
        ans = ""
        space_idx = 0
        for i, ch in enumerate(s):
            if space_idx < len(spaces) and i == spaces[space_idx]:
                ans += " "
                space_idx += 1
            ans += ch
        return ans
