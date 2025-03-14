class Solution:
    def isBalanced(self, num: str) -> bool:
        cnt = 0
        for i, val in enumerate(num):
            if i & 1 == 0:
                cnt += ord(val) - ord('0')
            else:
                cnt -= ord(val) - ord('0')
        return cnt == 0
