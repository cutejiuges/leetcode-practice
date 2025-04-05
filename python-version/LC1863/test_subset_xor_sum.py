from unittest import TestCase
from unittest import main
from .subset_xor_sum import Solution


class TestSolution(TestCase):
    def test_subset_xor_sum(self):
        solution = Solution()
        self.assertEqual(6, solution.subsetXORSum(nums=[1, 3]))
        self.assertEqual(28, solution.subsetXORSum(nums=[5, 1, 6]))
        self.assertEqual(480, solution.subsetXORSum(nums=[3, 4, 5, 6, 7, 8]))


if __name__ == '__main__':
    main()
