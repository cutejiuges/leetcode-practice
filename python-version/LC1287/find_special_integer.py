class Solution:
    def findSpecialInteger(self, arr: list[int]) -> int:
        n = len(arr)
        num, cnt = arr[0], 0
        for i in range(n):
            if arr[i] == num:
                cnt += 1
                if cnt * 4 > n:
                    return num
            else:
                num = arr[i]
                cnt = 1
        return -1