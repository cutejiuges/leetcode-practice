from unittest import TestCase
from .count_triples import Solution


class TestSolution(TestCase):
    def test_count_triples(self):
        solution = Solution()
        self.assertEqual(2, solution.countTriples(5))
        self.assertEqual(4, solution.countTriples(10))
