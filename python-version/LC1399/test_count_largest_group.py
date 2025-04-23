from unittest import TestCase
from .count_largest_group import Solution


class TestSolution(TestCase):
    def test_count_largest_group(self):
        solution = Solution()
        self.assertEqual(4, solution.countLargestGroup(n=13))
        self.assertEqual(2, solution.countLargestGroup(n=2))
        self.assertEqual(6, solution.countLargestGroup(n=15))
        self.assertEqual(5, solution.countLargestGroup(n=24))
