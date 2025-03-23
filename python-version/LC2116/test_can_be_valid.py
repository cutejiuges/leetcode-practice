import unittest
from .can_be_valid import Solution


class TestSolution(unittest.TestCase):
    def test_can_be_valid(self):
        solution = Solution()
        self.assertTrue(solution.canBeValid(s="))()))", locked="010100"))
        self.assertTrue(solution.canBeValid(s="()()", locked="0000"))
        self.assertFalse(solution.canBeValid(s=")", locked="0"))
        self.assertTrue(solution.canBeValid(s="(((())(((())", locked="111111010111"))


if __name__ == '__main__':
    unittest.main()
