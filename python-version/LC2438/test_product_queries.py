from unittest import TestCase
from .product_queries import Solution


class TestSolution(TestCase):
    def test_product_queries(self):
        solution = Solution()
        self.assertEqual([2, 4, 64], solution.productQueries(n=15, queries=[[0, 1], [2, 2], [0, 3]]))
        self.assertEqual([2], solution.productQueries(n=2, queries=[[0, 0]]))
