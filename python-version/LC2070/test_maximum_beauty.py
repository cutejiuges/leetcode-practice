import unittest
from .maximum_beauty import Solution


class TestSolution(unittest.TestCase):
    def test_maximum_beauty(self):
        solution = Solution()
        self.assertEqual([2, 4, 5, 5, 6, 6], solution.maximumBeauty(items=[[1, 2], [3, 2], [2, 4], [5, 6], [3, 5]],
                                                                    queries=[1, 2, 3, 4, 5, 6]))
        self.assertEqual([4], solution.maximumBeauty(items=[[1, 2], [1, 2], [1, 3], [1, 4]], queries=[1]))
        self.assertEqual([0], solution.maximumBeauty(items=[[10, 1000]], queries=[5]))


if __name__ == '__main__':
    unittest.main()
