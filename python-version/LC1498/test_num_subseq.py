from unittest import TestCase
from .num_subseq import Solution


class TestSolution(TestCase):
    def test_num_subseq(self):
        solution = Solution()
        self.assertEqual(4, solution.numSubseq(nums=[3, 5, 6, 7], target=9))
        self.assertEqual(6, solution.numSubseq(nums=[3, 3, 6, 8], target=10))
        self.assertEqual(61, solution.numSubseq(nums=[2, 3, 3, 4, 6, 7], target=12))
