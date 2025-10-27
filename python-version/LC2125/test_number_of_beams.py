from unittest import TestCase
from .number_of_beams import Solution


class TestSolution(TestCase):
    def test_number_of_beams(self):
        solution = Solution()
        self.assertEqual(8, solution.numberOfBeams(bank=["011001", "000000", "010100", "001000"]))
        self.assertEqual(0, solution.numberOfBeams(bank=["000", "111", "000"]))
