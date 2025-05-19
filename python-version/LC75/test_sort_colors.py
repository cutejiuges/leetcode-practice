from unittest import TestCase
from .sort_colors import Solution


class TestSolution(TestCase):
    def test_sort_colors(self):
        solution = Solution()
        nums = [2, 0, 2, 1, 1, 0]
        solution.sortColors(nums)
        self.assertEqual([0, 0, 1, 1, 2, 2], nums)
        nums = [2, 0, 1]
        solution.sortColors(nums)
        self.assertEqual([0, 1, 2], nums)
