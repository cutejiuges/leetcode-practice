class Solution:
    def maximumOr(self, nums: list[int], k: int) -> int:
        all_or, fixed = 0, 0
        for num in nums:
            fixed |= all_or & num  # 将固定位的1存放在fixed
            all_or |= num  # 将所有数字的or存放在all_or

        ans = 0
        for num in nums:
            ans = max(ans, (all_or ^ num) | fixed | (num << k))
            # 在all_or中去掉num的或，保留固定位的1,再将num左移k次或上去
        return ans
