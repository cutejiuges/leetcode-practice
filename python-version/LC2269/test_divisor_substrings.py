import unittest
from .divisor_substrings import Solution


class TestSolution(unittest.TestCase):
    def test_divisor_substrings(self):
        solution = Solution()
        self.assertEqual(2, solution.divisorSubstrings(num=240, k=2))
        self.assertEqual(2, solution.divisorSubstrings(num=430043, k=2))


if __name__ == '__main__':
    unittest.main()
