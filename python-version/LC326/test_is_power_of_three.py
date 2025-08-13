from unittest import TestCase
from .is_power_of_three import Solution


class TestSolution(TestCase):
    def test_is_power_of_three(self):
        solution = Solution()
        self.assertTrue(solution.isPowerOfThree(27))
        self.assertFalse(solution.isPowerOfThree(0))
        self.assertTrue(solution.isPowerOfThree(9))
        self.assertFalse(solution.isPowerOfThree(45))
