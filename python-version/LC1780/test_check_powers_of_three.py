from unittest import TestCase
from .check_powers_of_three import Solution


class TestSolution(TestCase):
    def test_check_powers_of_three(self):
        solution = Solution()
        self.assertTrue(solution.checkPowersOfThree(12))
        self.assertTrue(solution.checkPowersOfThree(91))
        self.assertFalse(solution.checkPowersOfThree(21))
