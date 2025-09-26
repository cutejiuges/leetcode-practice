from unittest import TestCase
from .triangle_number import Solution


class TestSolution(TestCase):
    def test_triangle_number(self):
        solution = Solution()
        self.assertEqual(3, solution.triangleNumber(nums=[2, 2, 3, 4]))
        self.assertEqual(4, solution.triangleNumber(nums=[4, 2, 3, 4]))
