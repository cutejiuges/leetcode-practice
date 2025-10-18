from unittest import TestCase
from .max_distinct_elements import Solution


class TestSolution(TestCase):
    def test_max_distinct_elements(self):
        solution = Solution()
        self.assertEqual(6, solution.maxDistinctElements(nums=[1, 2, 2, 3, 3, 4], k=2))
        self.assertEqual(3, solution.maxDistinctElements(nums=[4, 4, 4, 4], k=1))
