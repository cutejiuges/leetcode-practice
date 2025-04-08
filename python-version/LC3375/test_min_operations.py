from unittest import TestCase
from unittest import main
from .min_operations import Solution


class TestSolution(TestCase):
    def test_min_operations(self):
        solution = Solution()
        self.assertEqual(2, solution.minOperations(nums=[5, 2, 5, 4, 5], k=2))
        self.assertEqual(-1, solution.minOperations(nums=[2, 1, 2], k=2))
        self.assertEqual(4, solution.minOperations(nums=[9, 7, 5, 3], k=1))


if __name__ == '__main__':
    main()
