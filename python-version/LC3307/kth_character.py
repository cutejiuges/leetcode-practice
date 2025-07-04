class Solution:
    @staticmethod
    def kthCharacter(k: int, operations: list[int]) -> str:
        ans, t = 0, 0
        while k != 1:
            t = k.bit_length() - 1
            if 1 << t == k:
                t -= 1
            k -= (1 << t)
            if operations[t] == 1:
                ans += 1
        return chr(ord('a') + (ans % 26))
