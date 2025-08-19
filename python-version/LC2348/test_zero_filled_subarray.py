from unittest import TestCase
from .zero_filled_subarray import Solution


class TestSolution(TestCase):
    def test_zero_filled_subarray(self):
        solution = Solution()
        self.assertEqual(6, solution.zeroFilledSubarray([1, 3, 0, 0, 2, 0, 0, 4]))
        self.assertEqual(9, solution.zeroFilledSubarray([0, 0, 0, 2, 0, 0]))
        self.assertEqual(0, solution.zeroFilledSubarray([2, 10, 2019]))
