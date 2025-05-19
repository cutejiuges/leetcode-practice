class Solution:
    @staticmethod
    def triangleType(nums: list[int]) -> str:
        nums.sort()
        a, b, c = nums[0], nums[1], nums[2]
        if a + b <= c:
            return "none"
        if a == b == c:
            return "equilateral"
        if a == b or b == c or c == a:
            return "isosceles"
        return "scalene"
