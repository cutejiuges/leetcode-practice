from unittest import TestCase
from .maximum_length import Solution


class TestSolution(TestCase):
    def test_maximum_length(self):
        solution = Solution()
        self.assertEqual(4, solution.maximumLength(nums=[1, 2, 3, 4]))
        self.assertEqual(6, solution.maximumLength(nums=[1, 2, 1, 1, 2, 1, 2]))
        self.assertEqual(2, solution.maximumLength(nums=[1, 3]))
