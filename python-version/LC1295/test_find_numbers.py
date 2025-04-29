from unittest import TestCase
from .find_numbers import Solution


class TestSolution(TestCase):
    def test_find_numbers(self):
        solution = Solution()
        self.assertEqual(2, solution.findNumbers(nums=[12, 345, 2, 6, 7896]))
        self.assertEqual(1, solution.findNumbers(nums=[555, 901, 482, 1771]))
