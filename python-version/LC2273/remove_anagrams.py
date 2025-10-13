from typing import List


class Solution:
    def removeAnagrams(self, words: List[str]) -> List[str]:
        if not words:
            return []
        ans = [words[0]]
        idx = 0
        for word in words[1:]:
            if self.__anagram(word, ans[idx]):
                continue
            ans.append(word)
            idx += 1
        return ans

    @staticmethod
    def __anagram(a: str, b: str) -> bool:
        if len(a) != len(b):
            return False
        cnt = [0] * 26
        for c in a:
            cnt[ord(c) - ord('a')] += 1
        for c in b:
            cnt[ord(c) - ord('a')] -= 1
        for n in cnt:
            if n != 0:
                return False
        return True
