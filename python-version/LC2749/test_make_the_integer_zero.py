from unittest import TestCase
from .make_the_integer_zero import Solution


class TestSolution(TestCase):
    def test_make_the_integer_zero(self):
        solution = Solution()
        self.assertEqual(3, solution.makeTheIntegerZero(num1=3, num2=-2))
        self.assertEqual(-1, solution.makeTheIntegerZero(num1=5, num2=7))
