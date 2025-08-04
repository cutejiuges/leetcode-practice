from unittest import TestCase

from .total_fruit import Solution


class TestSolution(TestCase):
    def test_total_fruit(self):
        solution = Solution()
        self.assertEqual(3, solution.totalFruit(fruits=[1, 2, 1]))
        self.assertEqual(3, solution.totalFruit(fruits=[0, 1, 2, 2]))
        self.assertEqual(4, solution.totalFruit(fruits=[1, 2, 3, 2, 2]))
        self.assertEqual(5, solution.totalFruit(fruits=[3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4]))
