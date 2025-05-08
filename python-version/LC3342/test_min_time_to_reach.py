from unittest import TestCase
from .min_time_to_reach import Solution


class TestSolution(TestCase):
    def test_min_time_to_reach(self):
        solution = Solution()
        self.assertEqual(7, solution.minTimeToReach([[0, 4], [4, 4]]))
        self.assertEqual(6, solution.minTimeToReach([[0, 0, 0, 0], [0, 0, 0, 0]]))
        self.assertEqual(4, solution.minTimeToReach([[0, 1], [1, 2]]))
        self.assertEqual(71, solution.minTimeToReach([[0, 58], [27, 69]]))
