class Solution:
    @staticmethod
    def findLHS(nums: list[int]) -> int:
        ans = 0
        mp = dict()
        for num in nums:
            mp[num] = mp.get(num, 0) + 1

        for k, v in mp.items():
            if mp.__contains__(k+1):
                ans = max(ans, v + mp.get(k+1))
        return ans
