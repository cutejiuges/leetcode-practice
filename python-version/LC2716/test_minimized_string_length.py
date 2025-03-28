import unittest
from .minimized_string_length import Solution


class TestSolution(unittest.TestCase):
    def test_minimized_string_length(self):
        solution = Solution()
        self.assertEqual(3, solution.minimizedStringLength(s="aaabc"))
        self.assertEqual(3, solution.minimizedStringLength(s="cbbd"))
        self.assertEqual(2, solution.minimizedStringLength(s='aaaddd'))


if __name__ == '__main__':
    unittest.main()
