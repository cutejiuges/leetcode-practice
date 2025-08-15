from unittest import TestCase
from .is_power_of_four import Solution


class TestSolution(TestCase):
    def test_is_power_of_four(self):
        solution = Solution()
        self.assertTrue(solution.isPowerOfFour(4))
        self.assertFalse(solution.isPowerOfFour(5))
        self.assertTrue(solution.isPowerOfFour(1))
