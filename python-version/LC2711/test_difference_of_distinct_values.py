import unittest
from .difference_of_distinct_values import Solution


class TestSolution(unittest.TestCase):
    def test_difference_of_distinct_values(self):
        solution = Solution()
        self.assertEqual([[1, 1, 0], [1, 0, 1], [0, 1, 1]],
                         solution.differenceOfDistinctValues(grid=[[1, 2, 3], [3, 1, 5], [3, 2, 1]]))
        self.assertEqual([[0]], solution.differenceOfDistinctValues(grid=[[1]]))


if __name__ == '__main__':
    unittest.main()
