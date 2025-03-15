import unittest
from .score_of_string import Solution


class TestSolution(unittest.TestCase):
    def test_score_of_string(self):
        solution = Solution()
        self.assertEqual(13, solution.scoreOfString("hello"))
        self.assertEqual(50, solution.scoreOfString("zaz"))


if __name__ == '__main__':
    unittest.main()
