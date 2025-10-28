from unittest import TestCase
from .count_valid_selections import Solution


class TestSolution(TestCase):
    def test_count_valid_selections(self):
        solution = Solution()
        self.assertEqual(2, solution.countValidSelections(nums=[1, 0, 2, 0, 3]))
        self.assertEqual(0, solution.countValidSelections(nums=[2, 3, 4, 0, 4, 1, 0]))
