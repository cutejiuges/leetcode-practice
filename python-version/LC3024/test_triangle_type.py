from unittest import TestCase
from .triangle_type import Solution


class TestSolution(TestCase):
    def test_triangle_type(self):
        solution = Solution()
        self.assertEqual("equilateral", solution.triangleType(nums=[3, 3, 3]))
        self.assertEqual("scalene", solution.triangleType(nums=[3, 4, 5]))
