class Solution:
    def countOfSubstrings(self, word: str, k: int) -> int:
        return self.__calculate(word, k) - self.__calculate(word, k+1)

    @staticmethod
    def __calculate(ss: str, k: int) -> int:
        vowel_cnt_map = {}
        consonant_cnt = 0
        ans, low = 0, 0
        for ch in ss:
            if ch in "aeiou":
                vowel_cnt_map[ch] = vowel_cnt_map.get(ch, 0) + 1
            else:
                consonant_cnt += 1

            left_char = ss[low]
            while len(vowel_cnt_map) >= 5 and consonant_cnt >= k:
                if vowel_cnt_map.get(left_char, 0) > 0:
                    vowel_cnt_map[left_char] = vowel_cnt_map.get(left_char, 0) - 1
                    if vowel_cnt_map[left_char] == 0:
                        vowel_cnt_map.pop(left_char)
                else:
                    consonant_cnt -= 1
                low += 1
            ans += low
        return ans
