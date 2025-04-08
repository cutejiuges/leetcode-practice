class Solution:
    def minOperations(self, nums: list[int], k: int) -> int:
        st = set()
        for m in nums:
            if m < k:
                return -1
            elif m > k:
                st.add(m)
        return len(st)
