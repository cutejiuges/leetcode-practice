class Solution:
    def countCompleteSubarrays(self, nums: list[int]) -> int:
        left, ans = 0, 0
        distinct = self.__different_nums(nums)
        mp = dict()
        for i, _ in enumerate(nums):
            mp[nums[i]] = mp.get(nums[i], 0) + 1
            while len(mp) >= distinct:
                mp[nums[left]] = mp.get(nums[left]) - 1
                if mp[nums[left]] == 0:
                    del mp[nums[left]]
                left += 1
            ans += left
        return ans

    @staticmethod
    def __different_nums(nums: list[int]) -> int:
        st = set()
        for num in nums:
            st.add(num)
        return len(st)
