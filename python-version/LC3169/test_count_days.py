from unittest import TestCase
from .count_days import Solution


class TestSolution(TestCase):
    def test_count_days(self):
        solution = Solution()
        self.assertEqual(2, solution.countDays(days=10, meetings=[[5, 7], [1, 3], [9, 10]]))
        self.assertEqual(1, solution.countDays(days=5, meetings=[[2, 4], [1, 3]]))
        self.assertEqual(0, solution.countDays(days=6, meetings=[[1, 6]]))
