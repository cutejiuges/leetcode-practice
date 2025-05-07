from unittest import TestCase
from .min_time_to_reach import Solution


class TestSolution(TestCase):
    def test_min_time_to_reach(self):
        solution = Solution()
        self.assertEqual(6, solution.minTimeToReach([[0, 4], [4, 4]]))
        self.assertEqual(3, solution.minTimeToReach([[0, 0, 0], [0, 0, 0]]))
        self.assertEqual(3, solution.minTimeToReach([[0, 1], [1, 2]]))
