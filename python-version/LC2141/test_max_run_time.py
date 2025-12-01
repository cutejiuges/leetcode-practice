from unittest import TestCase
from .max_run_time import Solution


class TestSolution(TestCase):
    def test_max_run_time(self):
        solution = Solution()
        self.assertEqual(4, solution.maxRunTime(n=2, batteries=[3, 3, 3]))
        self.assertEqual(2, solution.maxRunTime(n=2, batteries=[1, 1, 1, 1]))
