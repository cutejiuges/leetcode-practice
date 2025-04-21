from unittest import TestCase
from .number_of_arrays import Solution


class TestSolution(TestCase):
    def test_number_of_arrays(self):
        solution = Solution()
        self.assertEqual(2, solution.numberOfArrays(differences=[1, -3, 4], lower=1, upper=6))
        self.assertEqual(4, solution.numberOfArrays(differences=[3, -4, 5, 1, -2], lower=-4, upper=5))
        self.assertEqual(0, solution.numberOfArrays(differences=[4, -7, 2], lower=3, upper=6))
