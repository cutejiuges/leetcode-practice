class Solution:
    @staticmethod
    def kthCharacter(k: int) -> str:
        ss = 'a'
        while k > len(ss):
            cur = ss
            for ch in cur:
                c = chr(ord('a') + ((ord(ch)-ord('a')+1) % 26))
                ss += c
        return ss[k-1]
