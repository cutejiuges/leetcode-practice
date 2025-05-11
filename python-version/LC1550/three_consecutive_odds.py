class Solution:
    @staticmethod
    def threeConsecutiveOdds(arr: list[int]) -> bool:
        for i in range(len(arr)-2):
            if arr[i] & 1 == 1 and arr[i+1] & 1 == 1 and arr[i+2] & 1 == 1:
                return True
        return False
