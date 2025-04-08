from unittest import TestCase
from unittest import main
from .minimum_operations import Solution


class TestSolution(TestCase):
    def test_minimum_operations(self):
        solution = Solution()
        self.assertEqual(2, solution.minimumOperations(nums=[1, 2, 3, 4, 2, 3, 3, 5, 7]))
        self.assertEqual(2, solution.minimumOperations(nums=[4, 5, 6, 4, 4]))
        self.assertEqual(0, solution.minimumOperations(nums=[6, 7, 8, 9]))


if __name__ == '__main__':
    main()
