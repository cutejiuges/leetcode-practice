from unittest import TestCase
from .min_sum import Solution


class TestSolution(TestCase):
    def test_min_sum(self):
        solution = Solution()
        self.assertEqual(12, solution.minSum(nums1=[3, 2, 0, 1, 0], nums2=[6, 5, 0]))
        self.assertEqual(-1, solution.minSum(nums1=[2, 0, 2, 0], nums2=[1, 4]))
