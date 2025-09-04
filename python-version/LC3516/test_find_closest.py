from unittest import TestCase
from .find_closest import Solution


class TestSolution(TestCase):
    def test_find_closest(self):
        solution = Solution()
        self.assertEqual(1, solution.findClosest(x=2, y=7, z=4))
        self.assertEqual(2, solution.findClosest(x=2, y=5, z=6))
        self.assertEqual(0, solution.findClosest(x=1, y=5, z=3))
