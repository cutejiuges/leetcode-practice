class Solution:
    @staticmethod
    def isValid(word: str) -> bool:
        if len(word) < 3:
            return False
        contains_vowel, contains_consonant = False, False
        for char in word:
            if char.isalpha():
                ch = char.lower()
                if ch in 'aeiou':
                    contains_vowel = True
                else:
                    contains_consonant = True
            elif not char.isdigit():
                return False
        return contains_vowel and contains_consonant
