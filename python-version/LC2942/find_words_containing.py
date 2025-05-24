class Solution:
    @staticmethod
    def findWordsContaining(words: list[str], x: str) -> list[int]:
        ans = []
        for i, word in enumerate(words):
            if x in word:
                ans.append(i)
        return ans
        