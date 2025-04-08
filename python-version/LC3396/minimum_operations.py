class Solution:
    def minimumOperations(self, nums: list[int]) -> int:
        st = set()
        n = len(nums)
        for i in range(n - 1, -1, -1):
            if nums[i] in st:
                return i // 3 + 1
            st.add(nums[i])
        return 0
