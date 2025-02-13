class Solution:
    def countBalls(self, lowLimit: int, highLimit: int) -> int:
        mp = {}
        mx = 0
        for ballNo in range(lowLimit, highLimit+1):
            curBall, boxNo = ballNo, 0
            while curBall > 0:
                boxNo += curBall % 10
                curBall //= 10
            cnt = mp.get(boxNo, 0)+1
            mp[boxNo] = cnt
            mx = max(mx, cnt)
        return mx