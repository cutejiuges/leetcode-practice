import unittest
from .count_prefixes import Solution


class TestSolution(unittest.TestCase):
    def test_count_prefixes(self):
        solution = Solution()
        self.assertEqual(3, solution.countPrefixes(words=["a", "b", "c", "ab", "bc", "abc"], s="abc"))
        self.assertEqual(2, solution.countPrefixes(words=["a", "a"], s="aa"))


if __name__ == '__main__':
    unittest.main()
