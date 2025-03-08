import unittest
from .maximum_beauty import Solution


class TestSolution(unittest.TestCase):
    def test_maximum_beauty(self):
        solution = Solution()
        self.assertEqual(14, solution.maximumBeauty(flowers=[1, 3, 1, 1], new_flowers=7, target=6, full=12, partial=1))
        self.assertEqual(30, solution.maximumBeauty(flowers=[2, 4, 5, 3], new_flowers=10, target=5, full=2, partial=6))


if __name__ == '__main__':
    unittest.main()
