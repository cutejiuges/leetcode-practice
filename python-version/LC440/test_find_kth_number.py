from unittest import TestCase
from .find_kth_number import Solution


class TestSolution(TestCase):
    def test_find_kth_number(self):
        solution = Solution()
        self.assertEqual(10, solution.findKthNumber(13, 2))
        self.assertEqual(1, solution.findKthNumber(1, 1))
        self.assertEqual(416126219, solution.findKthNumber(681692778,  351251360))
