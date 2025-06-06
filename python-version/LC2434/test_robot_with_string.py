from unittest import TestCase
from .robot_with_string import Solution


class TestSolution(TestCase):
    def test_robot_with_string(self):
        solution = Solution()
        self.assertEqual("azz", solution.robotWithString("zza"))
        self.assertEqual("abc", solution.robotWithString("bac"))
        self.assertEqual("addb", solution.robotWithString("bdda"))
