class Solution:
    def __init__(self):
        self.__visited = None

    def reorderedPowerOf2(self, n: int) -> bool:
        chars = str(n)
        self.__visited = [False] * len(chars)
        return self.__backtrack(chars, 0, 0)

    def __backtrack(self, chars: str, idx: int, num: int) -> bool:
        if idx == len(chars):
            return self.__powerOf2(num)
        for i in range(0, len(chars)):
            if chars[i] == '0' and num == 0:
                continue
            if self.__visited[i]:
                continue
            self.__visited[i] = True
            if self.__backtrack(chars, idx+1, num * 10 + ord(chars[i]) - ord('0')):
                return True
            self.__visited[i] = False
        return False

    @staticmethod
    def __powerOf2(n: int) -> bool:
        return n > 0 and (n & (n - 1)) == 0
