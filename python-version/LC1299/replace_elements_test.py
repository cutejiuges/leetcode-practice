import unittest
from replace_elements import Solution

class SolutionTest(unittest.TestCase):
    def test_replaceElements(self):
        solution = Solution()
        self.assertEqual([18,6,6,6,1,-1], solution.replaceElements(arr = [17,18,5,4,6,1]))
        self.assertEqual([-1], solution.replaceElements(arr = [400]))

if __name__ == '__main__':
    unittest.main()