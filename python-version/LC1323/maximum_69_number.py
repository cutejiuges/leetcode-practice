class Solution:
    @staticmethod
    def maximum69Number(num: int) -> int:
        ss = str(num)
        s = ss.replace('6', '9', 1)
        return int(s)
