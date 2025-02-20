import unittest
from even_odd_bit import Solution

class SolutionTest(unittest.TestCase):
    def test_evenOddBit(self):
        solution = Solution()
        self.assertEqual([1,2], solution.evenOddBit(50))
        self.assertEqual([0, 1], solution.evenOddBit(2))

if __name__ == '__main__':
    unittest.main()