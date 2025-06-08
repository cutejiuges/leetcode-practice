from unittest import TestCase
from .lexical_order import Solution


class TestSolution(TestCase):
    def test_lexical_order(self):
        solution = Solution()
        self.assertEqual([1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], solution.lexicalOrder(13))
        self.assertEqual([1, 2], solution.lexicalOrder(2))
