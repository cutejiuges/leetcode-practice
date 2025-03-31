import unittest

from .percentage_letter import Solution


class TestSolution(unittest.TestCase):
    def test_percentage_letter(self):
        solution = Solution()
        self.assertEqual(33, solution.percentageLetter(s="foobar", letter="o"))
        self.assertEqual(0, solution.percentageLetter(s="jjjj", letter='k'))


if __name__ == '__main__':
    unittest.main()
