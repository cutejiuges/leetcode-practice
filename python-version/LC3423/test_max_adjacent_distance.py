from unittest import TestCase
from .max_adjacent_distance import Solution


class TestSolution(TestCase):
    def test_max_adjacent_distance(self):
        solution = Solution()
        self.assertEqual(3, solution.maxAdjacentDistance(nums=[1, 2, 4]))
        self.assertEqual(5, solution.maxAdjacentDistance(nums=[-5, -10, -5]))
        self.assertEqual(3, solution.maxAdjacentDistance(nums=[-3, 0]))
