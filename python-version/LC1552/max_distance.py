class Solution:
    def maxDistance(self, position: list[int], m: int) -> int:
        position.sort()
        low, high = 1, position[-1]
        while low < high:
            mid = (low+high+1) // 2
            if self.__check(position, m, mid):
                low = mid
            else:
                high = mid-1
        return low
    
    def __check(self, position: list[int], m : int, mid : int):
        pre, cnt = position[0], 1
        for cur in position:
            if cur - pre >= mid:
                cnt += 1
                pre = cur
        return cnt >= m