import unittest
from find_special_integer import Solution

class SolutionTest(unittest.TestCase):
    def test_findSpecialInteger(self):
        solution = Solution()
        self.assertEqual(6, solution.findSpecialInteger(arr = [1,2,2,6,6,6,6,7,10]))

if __name__ == '__main__':
    unittest.main()