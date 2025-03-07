import unittest
from .beautiful_subsets import Solution


class TestSolution(unittest.TestCase):
    def test_beautiful_subsets(self):
        solution = Solution()
        self.assertEqual(4, solution.beautifulSubsets(nums=[2, 4, 6], k=2))
        self.assertEqual(1, solution.beautifulSubsets(nums=[1], k=1))


if __name__ == '__main__':
    unittest.main()
