import unittest
from .add_spaces import Solution


class TestSolution(unittest.TestCase):
    def test_add_spaces(self):
        solution = Solution()
        self.assertEqual("Leetcode Helps Me Learn", solution.addSpaces(s="LeetcodeHelpsMeLearn", spaces=[8, 13, 15]))
        self.assertEqual("i code in py thon", solution.addSpaces(s="icodeinpython", spaces=[1, 5, 7, 9]))
        self.assertEqual(" s p a c i n g", solution.addSpaces(s="spacing", spaces=[0, 1, 2, 3, 4, 5, 6]))


if __name__ == '__main__':
    unittest.main()
