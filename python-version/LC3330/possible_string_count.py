class Solution:
    @staticmethod
    def possibleStringCount(word: str) -> int:
        ans = 1
        length = len(word)
        for i in range(1, length):
            if word[i] == word[i-1]:
                ans += 1
        return ans
