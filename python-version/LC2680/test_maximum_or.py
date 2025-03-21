import unittest
from .maximum_or import Solution


class TestSolution(unittest.TestCase):
    def test_maximum_or(self):
        solution = Solution()
        self.assertEqual(30, solution.maximumOr(nums=[12, 9], k=1))
        self.assertEqual(35, solution.maximumOr(nums=[8, 1, 2], k=2))


if __name__ == '__main__':
    unittest.main()
