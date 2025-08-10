from unittest import TestCase
from .reordered_power_of_2 import Solution


class TestSolution(TestCase):
    def test_reordered_power_of2(self):
        solution = Solution()
        self.assertTrue(solution.reorderedPowerOf2(1))
        self.assertFalse(solution.reorderedPowerOf2(10))
