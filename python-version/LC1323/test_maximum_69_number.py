from unittest import TestCase
from .maximum_69_number import Solution


class TestSolution(TestCase):
    def test_maximum69number(self):
        solution = Solution()
        self.assertEqual(9969, solution.maximum69Number(9669))
        self.assertEqual(9999, solution.maximum69Number(9996))
        self.assertEqual(9999, solution.maximum69Number(9999))
