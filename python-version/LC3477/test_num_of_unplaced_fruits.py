from unittest import TestCase
from .num_of_unplaced_fruits import Solution


class TestSolution(TestCase):
    def test_num_of_unplaced_fruits(self):
        solution = Solution()
        self.assertEqual(1, solution.numOfUnplacedFruits(fruits=[4, 2, 5], baskets=[3, 5, 4]))
        self.assertEqual(0, solution.numOfUnplacedFruits(fruits=[3, 6, 1], baskets=[6, 4, 7]))
