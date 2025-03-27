import unittest
from .minimum_cost import Solution


class TestSolution(unittest.TestCase):
    def test_minimum_cost(self):
        solution = Solution()
        self.assertEqual(2, solution.minimumCost(s="0011"))
        self.assertEqual(9, solution.minimumCost(s="010101"))


if __name__ == '__main__':
    unittest.main()
