class Solution:
    @staticmethod
    def longestPalindrome(words: list[str]) -> int:
        mp = {}
        for word in words:
            mp[word] = mp.get(word, 0) + 1
        st = set()
        odd = False
        ans = 0

        for word, cnt in mp.items():
            reversed_word = word[::-1]
            if word == reversed_word:
                if cnt & 1 == 1:
                    odd = True
                ans += (cnt // 2) * 2 * 2
            elif word not in st:
                st.add(word)
                st.add(reversed_word)
                ans += min(mp.get(word, 0), mp.get(reversed_word, 0)) * 2 * 2
        if odd:
            ans += 2
        return ans
        