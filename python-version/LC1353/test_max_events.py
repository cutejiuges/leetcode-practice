from unittest import TestCase
from .max_events import Solution


class TestSolution(TestCase):
    def test_max_events(self):
        solution = Solution()
        self.assertEqual(3, solution.maxEvents([[1, 2], [2, 3], [3, 4]]))
        self.assertEqual(4, solution.maxEvents([[1, 2], [2, 3], [3, 4], [1, 2]]))
