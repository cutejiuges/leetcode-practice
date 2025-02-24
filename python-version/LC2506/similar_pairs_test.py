import unittest
from similar_pairs import Solution

class SolutionTest(unittest.TestCase):
    def test_silimarPairs(self):
        solution = Solution()
        self.assertEqual(2, solution.similarPairs(words = ["aba","aabb","abcd","bac","aabc"]))
        self.assertEqual(3, solution.similarPairs(words = ["aabb","ab","ba"]))
        self.assertEqual(0, solution.similarPairs(words = ["nba","cba","dba"]))

if __name__ == '__main__':
    unittest.main()