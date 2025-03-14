import unittest
from .is_balance import Solution


class TestSolution(unittest.TestCase):
    def test_is_balanced(self):
        solution = Solution()
        self.assertFalse(solution.isBalanced(num="1234"))
        self.assertTrue(solution.isBalanced(num="24123"))


if __name__ == '__main__':
    unittest.main()
