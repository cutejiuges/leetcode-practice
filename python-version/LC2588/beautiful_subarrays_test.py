import unittest
from beautiful_subarrays import Solution

class SolutionTest(unittest.TestCase):
    def test_beautifulSubarrays(self):
        solution = Solution()
        self.assertEqual(2, solution.beautifulSubarrays(nums = [4,3,1,2,4]))
        self.assertEqual(0, solution.beautifulSubarrays(nums = [1,10,4]))

if __name__ == '__main__':
    unittest.main()