class Solution:
    @staticmethod
    def sortVowels(s: str) -> str:
        vowel_set = {'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
        vowels = [ch for ch in s if ch in vowel_set]
        vowels.sort()
        ss = list(s)
        idx = 0
        for i, ch in enumerate(ss):
            if ch in vowel_set:
                ss[i] = vowels[idx]
                idx += 1
        return ''.join(ss)
