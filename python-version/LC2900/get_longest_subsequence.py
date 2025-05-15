class Solution:
    @staticmethod
    def getLongestSubsequence(words: list[str], groups: list[int]) -> list[str]:
        ans = list()
        n = len(groups)
        for i, _ in enumerate(groups):
            if i == n-1 or groups[i] != groups[i+1]:
                ans.append(words[i])
        return ans
